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
