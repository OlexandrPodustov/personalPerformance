#![allow(clippy::cargo_common_metadata)]
// LeetCode solution: `main` is an empty stub and the solution is exercised by
// the `#[cfg(test)]` cases below, so the non-test binary build sees it as dead.
#![allow(dead_code)]

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<Self>>,
}

impl ListNode {
    #[inline]
    const fn new(val: i32) -> Self {
        Self { next: None, val }
    }
}

pub struct Solution;

impl Solution {
    // l1/l2 are taken by value to match LeetCode's published method signature
    // (l2 is unused for now — solution is still a WIP stub).
    #[allow(clippy::needless_pass_by_value)]
    #[must_use]
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        if l1.is_none() && l2.is_none() {
            return None;
        }

        let mut current_l1 = l1;
        while let Some(node) = current_l1 {
            println!("current_l1: {} {:?}", node.val, node.next);
            current_l1 = node.next;
        }

        None
    }
}

const fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut head = Box::new(ListNode::new(1));
        head.next = Some(Box::new(ListNode::new(2)));
        let l1 = Some(head);
        let l2 = Some(Box::new(ListNode::new(2)));

        let result = Solution::add_two_numbers(l1, l2);
        let expected_result = Some(Box::new(ListNode::new(3)));
        assert_eq!(result, expected_result);
    }
}
