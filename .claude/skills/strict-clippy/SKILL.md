---
name: strict-clippy
description: Run cargo clippy across this workspace with all, pedantic, nursery, and cargo lint groups enabled, then fix the reported issues. Use when the user asks to lint, run clippy, or run "clippy-strict" on this repo.
---

# Strict clippy sweep (personalPerformance repo specifics)

This repo also has a **global** `strict-clippy` skill at
`~/.claude/skills/strict-clippy/SKILL.md` that holds everything generic: the
exact command, the JSON-dedupe triage method, the fix-vs-`#[allow]` philosophy,
and the gotchas (manifest `[lints]` losing to CLI group flags, workspace-glob
fragility, `cargo_common_metadata`/`multiple_crate_versions` judgment calls).
**Read that file first** — this one only covers what's specific to this repo, so
it doesn't duplicate (and can't drift out of sync with) the general guidance.

## What's specific to this repo

- Workspace root is `/Users/pod/go/src/github/personalPerformance/Cargo.toml`
  (run from anywhere under it; the alias `clippy-strict` in `~/.zshrc` works from
  any subdirectory). Members span `rust/`, `coursera/`, `exercism/`,
  `adventofcode/`, and `pandas/` — most of this is unrelated practice code, not
  just the LeetCode focus.
- The `rust/leetcode/*` and other glob members have already bitten us once (a
  stray `.claude/` dir broke `--workspace` resolution); it's excluded in the root
  `Cargo.toml`'s `exclude` list now. If `cargo clippy --workspace` errors with
  "failed to load manifest for workspace member", that's almost certainly this
  failure mode recurring — see the global skill for the fix.
- Before "fixing" a lint by completing solution logic, check whether the
  function is an intentional WIP stub (e.g. `add_two_numbers`, `minigrep::search`
  both still have stub bodies and failing tests on purpose) — only silence the
  lint (e.g. `#[allow(clippy::needless_pass_by_value)]` when a parameter is
  required by a fixed external signature like LeetCode's), don't implement the
  missing logic as a side effect of a lint sweep.
