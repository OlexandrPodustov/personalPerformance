---
name: strict-clippy
description: Run cargo clippy across the whole workspace with all, pedantic, nursery, and cargo lint groups enabled, then fix the reported issues. Use when the user asks to lint, run clippy, or run "clippy-strict" on this repo.
---

# Strict clippy sweep

Runs clippy with every major lint group turned on (warnings, not errors):

```bash
cargo clippy --workspace --all-targets -- -W clippy::all -W clippy::pedantic -W clippy::nursery -W clippy::cargo
```

A shell alias for the same command is available as `clippy-strict` (defined in `~/.zshrc`).

Run this from anywhere inside the workspace (cargo walks up to find the workspace root
`Cargo.toml` at the repo root, which lists members across `rust/`, `coursera/`,
`exercism/`, `adventofcode/`, and `pandas/`).

## Fixing issues

- Output has massive duplication: the same ~83 warnings repeat per target (lib/bin/test
  variants of the same crate). Use `--message-format=json` and dedupe by
  `(file, line_start, lint code)` before triaging — otherwise it looks like hundreds of
  distinct problems when it's usually a few dozen.
- Prefer fixing the lint at the source over `#[allow(...)]`, unless the lint is fighting
  a deliberate choice (e.g. an intentionally unfinished learning stub — see project notes
  on WIP exercise stubs before "completing" any LeetCode solution logic just to silence
  a lint).
- Common recurring lints in this repo: `uninlined_format_args` (inline the variable into
  the `format!`/`println!` string), `use_self` (use `Self` instead of repeating the
  struct name), `missing_const_for_fn` (mark trivial constructors `const fn`),
  `must_use_candidate` (add `#[must_use]` to pure accessor/computation methods),
  `needless_pass_by_value`, `cast_possible_truncation`/`cast_possible_wrap` (use
  `try_from`/`u32`-sized types instead of `as` casts where it's a real risk, e.g. indices
  that could exceed `i32`).
- `clippy::cargo`'s `multiple_crate_versions` warning (e.g. duplicate `hashbrown` or
  `wit-bindgen` versions) is a dependency-graph issue, not a code issue — leave it unless
  asked to chase down `Cargo.lock` duplication.

## Gotchas learned the hard way

- **Manifest `[lints]` tables don't reliably beat CLI group flags.** Allowing a specific
  lint via `[workspace.lints.clippy]` (with member crates opting in via
  `lints.workspace = true`) looks like the "proper" cargo-native way to suppress it, but
  cargo translates that into a CLI flag too — and when this command's own
  `-W clippy::cargo` group flag is appended after it, the group re-enables the
  specifically-allowed lint. Only a source-level `#![allow(clippy::lint_name)]` attribute
  at the crate root reliably wins over a CLI group flag. Don't waste time on the manifest
  approach for this command; go straight to the attribute.
- **Workspace member globs are fragile.** `rust/leetcode/*` (and similar globs in this
  repo's root `Cargo.toml`) break `cargo` entirely if any directory matching the glob
  lacks a `Cargo.toml` — including stray dotfiles/dirs the Claude Code harness can drop
  while working in a subfolder (e.g. a per-directory `.claude/settings.local.json`). If
  `cargo clippy --workspace` suddenly errors with "failed to load manifest for workspace
  member", check for a new stray directory and add it to `exclude` in the root
  `Cargo.toml` rather than deleting it blindly (it may hold real permission grants).
