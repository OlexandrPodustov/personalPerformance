# Project instructions

## Working style

- **Ask questions first — never assume.** Use the `AskUserQuestion` tool whenever
  the request is ambiguous or two readings would lead to materially different work.

## Tool-permission model (reference)

Design notes for a tiered auto-approval pipeline. Each tier either decides or
falls through to the next; anything that reaches the end prompts the user
normally.

### Tier 1 — MCP annotations (~0ms)

Trust the tool's own declared hints:

| Hint              | Decision |
| ----------------- | -------- |
| `readOnlyHint`    | approve  |
| `destructiveHint` | deny     |

### Tier 2 — SHA-256 cache (~0ms)

Key on `sha256(tool + cmd)`. A cache hit replays the previous decision, so an
identical call is never re-judged.

### Tier 3 — LLM-as-a-judge (~2–18s)

`verdict = llm.classify(cmd, SAFETY_PROMPT)`

| Verdict  | Examples                                      |
| -------- | --------------------------------------------- |
| `SAFE`   | `cat`, `ls`, `git status`, `npm test`, `kubectl get` |
| `UNSAFE` | `rm -rf`, `terraform apply`, `sudo`, `env \| SECRET` |

**Fail-safe:** any error → `exit(0)` → fall back to the normal permission prompt.
The classifier can only ever _skip_ a prompt, never suppress one it failed to
evaluate.

### Classifier policy

A Sonnet classifier evaluates every tool call.

- **Allows:** file read/write, `npm install`, push to `feature/*`, read-only HTTP
- **Blocks:** download-and-execute, credential leaks, force-push to `main`, mass deletion

**Circuit breaker:** 3 consecutive blocks → stop auto-deciding and resume prompting.

### Limits of the deterministic rules

The `permissions.deny` list in `.claude/settings.json` is a fast first pass, not a
security boundary. Know what it cannot do:

- **Prefix matching only.** `Bash(git push --force *)` catches
  `git push --force origin main` but not `git push origin main --force` — flag
  order defeats it. The same applies to every `Bash(...)` rule.
- **`Read(...)` denies gate the Read tool, not the shell.** A denied
  `Read(//**/.env)` does nothing to stop `cat .env` in Bash.
- **Compound and indirect forms slip through.** `curl -o f url && sh f` is not
  the denied `curl * | sh`.

Tier 3 is what actually holds: the classifier reads intent rather than matching
strings, and `autoMode.hard_deny` is where the real prohibitions live. Treat the
deny list as defense in depth, never as the only fence.

### What belongs in which file

`.claude/settings.json` is committed to a **public** repo — anyone who clones it
inherits those rules. Keep it to non-executing operations. Anything that compiles
or runs repository code (`cargo build`/`test`/`clippy`, `go build`/`test`, `npm
install`) executes `build.rs`, proc macros, test bodies, and install scripts, so
it belongs in the git-ignored `.claude/settings.local.json` alongside
machine-specific paths and the `defaultMode: "auto"` opt-in.
