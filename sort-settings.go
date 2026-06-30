// sort-settings: alphabetically sort the top-level keys of a JSONC settings
// file (e.g. VS Code settings.json), preserving comments, nested-object order,
// and trailing commas. Each comment block stays attached to the key that
// follows it. Parsed values are verified unchanged (deep-equal) before writing.
//
// Build:
//
//	go build -o sort-settings.exe sort-settings.go
//
// Usage (PowerShell):
//
//	.\sort-settings.exe path\to\settings.json
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// analyze returns the index just past the root '{', the index of the matching
// '}', and the indices of every comma at the top level (depth 1). Strings and
// comments are skipped so braces/commas inside them do not affect depth.
func analyze(s string) (bodyStart, rootClose int, commas []int) {
	root := strings.IndexByte(s, '{')
	i := root + 1
	depth := 1
	n := len(s)
	for i < n {
		c := s[i]
		switch {
		case c == '"':
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
			continue
		case c == '/' && i+1 < n && s[i+1] == '/':
			if k := strings.IndexByte(s[i:], '\n'); k == -1 {
				i = n
			} else {
				i += k
			}
			continue
		case c == '/' && i+1 < n && s[i+1] == '*':
			if k := strings.Index(s[i:], "*/"); k == -1 {
				i = n
			} else {
				i += k + 2
			}
			continue
		case c == '{' || c == '[':
			depth++
		case c == '}' || c == ']':
			depth--
			if depth == 0 {
				return root + 1, i, commas
			}
		case c == ',' && depth == 1:
			commas = append(commas, i)
		}
		i++
	}
	return root + 1, n, commas
}

// findKey returns the content of the first JSON string in seg, skipping any
// leading comments. That string is the property key of the segment.
func findKey(seg string) string {
	i, n := 0, len(seg)
	for i < n {
		c := seg[i]
		switch {
		case c == '"':
			var b strings.Builder
			j := i + 1
			for j < n {
				if seg[j] == '\\' {
					if j+1 < n {
						b.WriteByte(seg[j+1])
					}
					j += 2
					continue
				}
				if seg[j] == '"' {
					break
				}
				b.WriteByte(seg[j])
				j++
			}
			return b.String()
		case c == '/' && i+1 < n && seg[i+1] == '/':
			if k := strings.IndexByte(seg[i:], '\n'); k == -1 {
				return ""
			} else {
				i += k + 1
			}
			continue
		case c == '/' && i+1 < n && seg[i+1] == '*':
			if k := strings.Index(seg[i:], "*/"); k == -1 {
				return ""
			} else {
				i += k + 2
			}
			continue
		}
		i++
	}
	return ""
}

// trimBlank drops blank lines from the start and end of seg, keeping the
// indentation of the remaining content lines intact.
func trimBlank(seg string) string {
	lines := strings.Split(seg, "\n")
	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	end := len(lines)
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}
	return strings.Join(lines[start:end], "\n")
}

var trailingComma = regexp.MustCompile(`,(\s*[}\]])`)

// stripJSONC removes comments (string-aware) and trailing commas so the result
// can be parsed by a strict JSON decoder.
func stripJSONC(s string) string {
	var b strings.Builder
	i, n := 0, len(s)
	for i < n {
		c := s[i]
		switch {
		case c == '"':
			b.WriteByte(c)
			i++
			for i < n {
				b.WriteByte(s[i])
				if s[i] == '\\' {
					if i+1 < n {
						b.WriteByte(s[i+1])
					}
					i += 2
					continue
				}
				if s[i] == '"' {
					i++
					break
				}
				i++
			}
			continue
		case c == '/' && i+1 < n && s[i+1] == '/':
			if k := strings.IndexByte(s[i:], '\n'); k == -1 {
				i = n
			} else {
				i += k
			}
			continue
		case c == '/' && i+1 < n && s[i+1] == '*':
			if k := strings.Index(s[i:], "*/"); k == -1 {
				i = n
			} else {
				i += k + 2
			}
			continue
		}
		b.WriteByte(c)
		i++
	}
	return trailingComma.ReplaceAllString(b.String(), "$1")
}

type entry struct {
	key string
	seg string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: sort-settings <file.json>")
		os.Exit(2)
	}
	path := os.Args[1]
	raw, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading:", err)
		os.Exit(2)
	}
	nl := "\n"
	if strings.Contains(string(raw), "\r\n") {
		nl = "\r\n"
	}
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")

	bodyStart, rootClose, commas := analyze(text)
	bounds := []int{bodyStart}
	ends := []int{}
	for _, c := range commas {
		bounds = append(bounds, c+1)
		ends = append(ends, c)
	}
	ends = append(ends, rootClose)

	var entries []entry
	for idx := range bounds {
		seg := trimBlank(text[bounds[idx]:ends[idx]])
		key := findKey(seg)
		if key == "" {
			continue // empty segment left by a trailing comma
		}
		entries = append(entries, entry{key, seg})
	}
	sort.SliceStable(entries, func(a, b int) bool {
		return strings.ToLower(entries[a].key) < strings.ToLower(entries[b].key)
	})

	var out strings.Builder
	out.WriteString("{\n")
	for idx, e := range entries {
		out.WriteString(e.seg)
		if idx < len(entries)-1 {
			out.WriteByte(',')
		}
		out.WriteByte('\n')
	}
	out.WriteString("}\n")
	result := out.String()

	var origObj, newObj map[string]any
	if err := json.Unmarshal([]byte(stripJSONC(text)), &origObj); err != nil {
		fmt.Fprintln(os.Stderr, "cannot parse original:", err)
		os.Exit(1)
	}
	if err := json.Unmarshal([]byte(stripJSONC(result)), &newObj); err != nil {
		fmt.Fprintln(os.Stderr, "cannot parse result:", err)
		os.Exit(1)
	}
	if !reflect.DeepEqual(origObj, newObj) {
		fmt.Fprintln(os.Stderr, "values changed; aborting without writing")
		os.Exit(1)
	}

	if err := os.WriteFile(path, []byte(strings.ReplaceAll(result, "\n", nl)), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "error writing:", err)
		os.Exit(1)
	}
	fmt.Printf("OK: %d keys, sorted, values identical\n", len(newObj))
}
