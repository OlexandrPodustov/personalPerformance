# Singly Linked List — Rust Cheatsheet

```rust
pub struct ListNode { pub val: i32, pub next: Option<Box<ListNode>> }
```

Key idea: every "slot" — the `head` and every `.next` — has the same type
`Option<Box<ListNode>>`, so one `&mut` cursor can walk them all.

## Build front-to-back (preserve order) — tail cursor
```rust
let mut head: Option<Box<ListNode>> = None;
let mut tail = &mut head;                 // cursor on the trailing empty slot
for v in source {
    *tail = Some(Box::new(ListNode { val: v, next: None })); // fill
    if let Some(node) = tail {
        tail = &mut node.next;            // follow
    }
}
head
```
Invariant: after each step, `tail` points at the empty `.next` at the end.

## Build back-to-front (reverses order) — prepend
```rust
let mut head: Option<Box<ListNode>> = None;
for v in source {
    head = Some(Box::new(ListNode { val: v, next: head.take() }));
}
head
```

## Reverse a list — move nodes one by one
```rust
let mut prev = None;
let mut cur = head;
while let Some(mut node) = cur {
    cur = node.next.take();   // detach rest
    node.next = prev;         // point backwards
    prev = Some(node);
}
prev
```

## Traverse while consuming (read each val)
```rust
let mut cur = head;
while let Some(node) = cur {
    // use node.val
    cur = node.next;          // move ownership forward
}
```

## Traverse by borrowing (keep the list)
```rust
let mut cur = &head;
while let Some(node) = cur {
    // use node.val
    cur = &node.next;
}
```

Notes:
- `*tail = ...` writes through the reference into the slot; without `*` you reassign the cursor.
- `if let Some(node) = tail` reborrows; `&mut node.next` is again `&mut Option<Box<ListNode>>` — avoids `unwrap`.
- `take()` swaps a value out leaving `None`, sidestepping move/borrow conflicts.
