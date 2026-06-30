// rstest-fmt: format the argument lists of rstest attribute macros that
// `cargo fmt` leaves untouched (e.g. `#[case(&[2,4,3], &[5,6,4])]`).
//
// rustfmt never reformats the body of an attribute macro. But the body of
// `#[case(...)]` is just a comma-separated list of normal Rust expressions,
// so for each matching attribute we splice its arguments into a throwaway
// function call, hand THAT to rustfmt, and read the normalized arguments
// back out. The result is genuine rustfmt spacing, not a regex guess.
//
// Build:
//
//	go build -o rstest-fmt.exe rstest-fmt.go
//
// Usage (PowerShell):
//
//	.\rstest-fmt.exe src\foo.rs                       # rewrite in place
//	.\rstest-fmt.exe -check (git ls-files '*.rs')     # exit 1 if changes needed
//	.\rstest-fmt.exe -stdout src\foo.rs               # print result, don't write
//	git ls-files '*.rs' | .\rstest-fmt.exe -          # read paths from stdin
//	Get-Content foo.rs | .\rstest-fmt.exe -stdin      # rustfmt + attrs, source in/out (editor filter)
//
// Flags must come before file arguments (standard Go flag parsing).
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' || c == '\v'
}

func isIdentOrColon(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '_' || c == ':'
}

// skipStringOrComment reports whether s[i] begins a string/char/comment and,
// if so, returns the index just past it.
func skipStringOrComment(s []byte, i int) (int, bool) {
	n := len(s)
	c := s[i]

	// line comment
	if c == '/' && i+1 < n && s[i+1] == '/' {
		if j := bytes.IndexByte(s[i:], '\n'); j != -1 {
			return i + j, true
		}
		return n, true
	}
	// block comment (nested)
	if c == '/' && i+1 < n && s[i+1] == '*' {
		depth := 1
		j := i + 2
		for j < n && depth > 0 {
			switch {
			case s[j] == '/' && j+1 < n && s[j+1] == '*':
				depth++
				j += 2
			case s[j] == '*' && j+1 < n && s[j+1] == '/':
				depth--
				j += 2
			default:
				j++
			}
		}
		return j, true
	}
	// raw string: r"..."  r#"..."#  br#"..."#
	if c == 'r' || c == 'b' {
		k := i
		if s[k] == 'b' {
			k++
		}
		if k < n && s[k] == 'r' {
			k++
			hashes := 0
			for k < n && s[k] == '#' {
				hashes++
				k++
			}
			if k < n && s[k] == '"' {
				closing := append([]byte{'"'}, bytes.Repeat([]byte{'#'}, hashes)...)
				if idx := bytes.Index(s[k+1:], closing); idx != -1 {
					return k + 1 + idx + len(closing), true
				}
				return n, true
			}
		}
		return 0, false
	}
	// normal string
	if c == '"' {
		j := i + 1
		for j < n {
			if s[j] == '\\' {
				j += 2
				continue
			}
			if s[j] == '"' {
				return j + 1, true
			}
			j++
		}
		return n, true
	}
	// char literal vs lifetime: only skip if it actually closes with '
	if c == '\'' {
		j := i + 1
		if j < n && s[j] == '\\' {
			j += 2
			if j < n && s[j] == '\'' {
				return j + 1, true
			}
			return 0, false
		}
		if j+1 < n && s[j+1] == '\'' {
			return j + 2, true
		}
		return 0, false // a lifetime, not a char literal
	}
	return 0, false
}

type span struct{ start, end int }

// findAttrArgs returns the spans of text inside the outer parens of every
// `#[<attr>...(...)]` whose first path segment is in attrs.
func findAttrArgs(s []byte, attrs map[string]bool) []span {
	n := len(s)
	var spans []span
	i := 0
	for i < n {
		if ni, ok := skipStringOrComment(s, i); ok && ni > i {
			i = ni
			continue
		}
		if s[i] == '#' && i+1 < n &&
			(s[i+1] == '[' || (s[i+1] == '!' && i+2 < n && s[i+2] == '[')) {
			j := i + 2
			if s[i+1] == '!' {
				j = i + 3
			}
			for j < n && isSpace(s[j]) {
				j++
			}
			startPath := j
			for j < n && isIdentOrColon(s[j]) {
				j++
			}
			path := string(s[startPath:j])
			firstSeg := path
			if idx := strings.Index(path, "::"); idx >= 0 {
				firstSeg = path[:idx]
			}
			for j < n && isSpace(s[j]) {
				j++
			}
			if attrs[firstSeg] && j < n && s[j] == '(' {
				innerStart := j + 1
				depth := 0
				k := j
				for k < n {
					if nk, ok := skipStringOrComment(s, k); ok && nk > k {
						k = nk
						continue
					}
					switch s[k] {
					case '(':
						depth++
					case ')':
						depth--
						if depth == 0 {
							spans = append(spans, span{innerStart, k})
						}
					}
					if depth == 0 && s[k] == ')' {
						break
					}
					k++
				}
				i = k + 1
				continue
			}
		}
		i++
	}
	return spans
}

// rustfmtArgs round-trips a `#[case(...)]` argument list through rustfmt.
func rustfmtArgs(argText string, rustfmtCmd []string, edition string) string {
	wrapped := "fn __rstest_fmt() {\n    __m(" + argText + ");\n}\n"
	args := append([]string{}, rustfmtCmd[1:]...)
	args = append(args, "--edition", edition, "--config", "max_width=100000")
	cmd := exec.Command(rustfmtCmd[0], args...)
	cmd.Stdin = strings.NewReader(wrapped)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return argText // doesn't parse standalone: leave untouched
	}
	o := out.String()
	openIdx := strings.Index(o, "__m(")
	if openIdx == -1 {
		return argText
	}
	p := openIdx + len("__m(")
	start := p
	depth := 1
	for p < len(o) {
		if o[p] == '(' {
			depth++
		} else if o[p] == ')' {
			depth--
			if depth == 0 {
				break
			}
		}
		p++
	}
	inner := o[start:p]
	// A line comment cannot survive being folded onto one line — everything
	// after "//" would become a comment. Return the original text unchanged.
	if hasLineComment(inner) {
		return argText
	}
	// fold to one line: rustfmt won't wrap at max_width=100000, but be safe
	lines := strings.Split(inner, "\n")
	for idx := range lines {
		lines[idx] = strings.TrimSpace(lines[idx])
	}
	return strings.TrimSpace(strings.Join(lines, " "))
}

// hasLineComment reports whether s contains a "//" line comment outside of
// string literals.
func hasLineComment(s string) bool {
	i, n := 0, len(s)
	for i < n {
		switch s[i] {
		case '"':
			i++
			for i < n {
				if s[i] == '\\' {
					i += 2
					continue
				}
				if s[i] == '"' {
					i++
					break
				}
				i++
			}
		case '/':
			if i+1 < n && s[i+1] == '/' {
				return true
			}
			i++
		default:
			i++
		}
	}
	return false
}

func formatSource(src []byte, rustfmtCmd []string, edition string, attrs map[string]bool) []byte {
	spans := findAttrArgs(src, attrs)
	// reverse order keeps earlier indices valid after each splice
	for idx := len(spans) - 1; idx >= 0; idx-- {
		is, ie := spans[idx].start, spans[idx].end
		original := string(src[is:ie])
		formatted := rustfmtArgs(original, rustfmtCmd, edition)
		if formatted != original {
			nb := make([]byte, 0, len(src)+len(formatted)-len(original))
			nb = append(nb, src[:is]...)
			nb = append(nb, formatted...)
			nb = append(nb, src[ie:]...)
			src = nb
		}
	}
	return src
}

// rustfmtWhole runs rustfmt over an entire source buffer, returning the input
// unchanged if rustfmt rejects it (e.g. a syntax error during an active edit).
func rustfmtWhole(src []byte, rustfmtCmd []string, edition string) []byte {
	args := append([]string{}, rustfmtCmd[1:]...)
	args = append(args, "--edition", edition)
	cmd := exec.Command(rustfmtCmd[0], args...)
	cmd.Stdin = bytes.NewReader(src)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return src
	}
	return out.Bytes()
}

func main() {
	attrsFlag := flag.String("attrs", "case,values", "comma-separated attribute names to format")
	edition := flag.String("edition", "2024", "edition passed to rustfmt")
	rustfmtFlag := flag.String("rustfmt", "rustfmt", `rustfmt command, e.g. "rustup run nightly rustfmt"`)
	check := flag.Bool("check", false, "do not write; exit 1 if any file would change")
	stdout := flag.Bool("stdout", false, "print result instead of writing the file")
	stdin := flag.Bool("stdin", false, "read source from stdin, run rustfmt then format attrs, write to stdout")
	flag.Parse()

	attrs := map[string]bool{}
	for _, a := range strings.Split(*attrsFlag, ",") {
		if a = strings.TrimSpace(a); a != "" {
			attrs[a] = true
		}
	}
	rustfmtCmd := strings.Fields(*rustfmtFlag)
	if len(rustfmtCmd) == 0 {
		rustfmtCmd = []string{"rustfmt"}
	}

	// editor filter mode: format a whole buffer (rustfmt + attr args) from stdin to stdout
	if *stdin {
		src, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading stdin:", err)
			os.Exit(2)
		}
		out := formatSource(rustfmtWhole(src, rustfmtCmd, *edition), rustfmtCmd, *edition, attrs)
		os.Stdout.Write(out)
		return
	}

	files := flag.Args()
	// support `-` to read newline-separated paths from stdin (handy with git ls-files)
	if len(files) == 1 && files[0] == "-" {
		files = files[:0]
		sc := bufio.NewScanner(os.Stdin)
		sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
		for sc.Scan() {
			if line := strings.TrimSpace(sc.Text()); line != "" {
				files = append(files, line)
			}
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading stdin:", err)
			os.Exit(2)
		}
	}
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "usage: rstest-fmt [flags] file.rs ...   (or '-' to read paths from stdin)")
		flag.PrintDefaults()
		os.Exit(2)
	}

	changedAny := false
	exitCode := 0
	for _, path := range files {
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", path, err)
			exitCode = 2
			continue
		}
		out := formatSource(src, rustfmtCmd, *edition, attrs)
		if !bytes.Equal(out, src) {
			changedAny = true
			switch {
			case *check:
				fmt.Fprintf(os.Stderr, "would reformat: %s\n", path)
			case *stdout:
				os.Stdout.Write(out)
			default:
				info, _ := os.Stat(path)
				mode := os.FileMode(0644)
				if info != nil {
					mode = info.Mode()
				}
				if err := os.WriteFile(path, out, mode); err != nil {
					fmt.Fprintf(os.Stderr, "error writing %s: %v\n", path, err)
					exitCode = 2
					continue
				}
				fmt.Fprintf(os.Stderr, "reformatted: %s\n", path)
			}
		} else if *stdout {
			os.Stdout.Write(out)
		}
	}

	if *check && changedAny && exitCode == 0 {
		exitCode = 1
	}
	os.Exit(exitCode)
}
