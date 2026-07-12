// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::cell::RefCell;
use std::rc::Rc;

pub struct Solution;

impl Solution {
    pub fn has_path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {
        _ = root;
        _ = target_sum;

        false
    }
}
#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;
    use std::collections::VecDeque;

    fn to_tree(list: Vec<Option<i32>>) -> Option<Rc<RefCell<TreeNode>>> {
        if list.is_empty() {
            return None;
        }

        let tree = TreeNode::new(list[0].unwrap());
        let mut parents: VecDeque<Option<i32>> = VecDeque::new();
        for i in 0..list.len() {
            parents.push_back(list[i]);
        }

        Some(Rc::new(RefCell::new(tree)))
    }

    #[rstest]
    // #[case(vec![5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1], 22, true)]
    #[case(None, 22, false)]
    fn it_works(
        #[case] root: Option<Rc<RefCell<TreeNode>>>,
        #[case] k: i32,
        #[case] expected_result: bool,
    ) {
        let res = to_tree(vec![Some(1), Some(2), Some(3), None]);
        println!("res: {:?}", res);

        let res = Solution::has_path_sum(root, k);
        assert_eq!(res, expected_result);
    }
}
