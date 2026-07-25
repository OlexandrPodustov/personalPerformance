# LeetCode review log (spaced repetition)

**How this works** (adaptive expanding intervals):

- A **review** = re-solve the problem in Rust from a blank file, cold, narrating aloud in English. ~10–20 min per easy, ~25–35 per medium.
- After first solving a problem, set `next_review` ≈ +3 days. After each **clean** cold re-solve, roughly double the interval: 3 → 7 → 14 → 30 → 90 days. Struggled → halve it or repeat tomorrow. Clean at ~90 days = the pattern is yours; drop it from active rotation.
- Solved unaided on the first try in a familiar pattern → may start directly at +30d.
- Do everything due during the weekly review session; update `next_review` and `notes` after each pass.

| LC  | problem                             | pattern                                     | solved     | unaided | next_review | notes                                               |
| --- | ----------------------------------- | ------------------------------------------- | ---------- | ------- | ----------- | --------------------------------------------------- |
| 1   | two_sum                             | hash map, one-pass complement               | 2026-06-12 | y       | 2026-07-26  | seeded retroactively from git                       |
| 2   | add_two_numbers                     | linked list + carry (Box, take/tail-cursor) | 2026-06-12 | y       | 2026-08-02  | the valuable re-solve: Box ownership                |
| 125 | valid-palindrome                    | two pointers (iterator .eq(rev))            | 2026-06-16 | y       | 2026-08-09  | try the index-based half-scan variant on review     |
| 27  | remove-element                      | two pointers, overwrite-keep                | 2026-06-21 | y       | 2026-08-16  | on review: also strengthen the test (assert prefix) |
| 88  | merge-sorted-array                  | three-cursor reverse merge                  | 2026-07-06 | y       | 2026-08-05  |                                                     |
| 26  | remove-duplicates-from-sorted-array | two pointers, read/write dedup              | 2026-07-11 | y       | 2026-08-10  | on review: assert prefix contents in test           |
| 112 | path-sum                            | tree DFS root-to-leaf (Rc<RefCell>)         | —          | —       | in progress | stub since 2026-07-12; Week 1 of Phase A            |
