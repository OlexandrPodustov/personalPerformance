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
    // l1/l2 are taken by value to match LeetCode's published method signature.
    #[allow(clippy::needless_pass_by_value)]
    #[must_use]
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        fn pop_front(list: &mut Option<Box<ListNode>>) -> i32 {
            list.take().map_or(0, |n| {
                *list = n.next;
                n.val
            })
        }

        if l1.is_none() && l2.is_none() {
            return None;
        }

        let mut head: Option<Box<ListNode>> = None;
        let mut tail = &mut head;

        let mut l1 = l1;
        let mut l2 = l2;
        let mut remainder = 0;

        while l1.is_some() || l2.is_some() || remainder > 0 {
            let sum = remainder + pop_front(&mut l1) + pop_front(&mut l2);
            remainder = sum / 10;

            *tail = Some(Box::new(ListNode::new(sum % 10)));
            tail = &mut tail.as_mut().unwrap().next;
        }

        head
    }
}

const fn main() {}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    impl ListNode {
        fn list(vals: &[i32]) -> Option<Box<ListNode>> {
            vals.iter()
                .rev()
                .fold(None, |next, &val| Some(Box::new(ListNode { val, next })))
        }
    }

    #[rstest]
    #[case(&[1, 2, 3], &[5, 5, 5], &[6, 7, 8])]
    #[case(&[0], &[0], &[0])]
    #[case(&[9, 9, 9], &[1], &[0, 0, 0, 1])]
    #[case(&[2, 4, 3], &[5, 6, 4], &[7, 0, 8])]
    #[case(&[9, 9, 9, 9, 9, 9, 9], &[9, 9, 9, 9], &[8, 9, 9, 9, 0, 0, 0, 1])]
    fn it_works(#[case] l1_vals: &[i32], #[case] l2_vals: &[i32], #[case] expected_vals: &[i32]) {
        let l1 = ListNode::list(l1_vals);
        let l2 = ListNode::list(l2_vals);
        let result = Solution::add_two_numbers(l1, l2);
        let expected_result = ListNode::list(expected_vals);
        assert_eq!(result, expected_result);
    }
}
