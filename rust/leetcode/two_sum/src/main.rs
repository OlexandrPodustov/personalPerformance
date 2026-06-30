// LeetCode solution: `main` is an empty stub and the solution is exercised by
// the `#[cfg(test)]` cases below, so the non-test binary build sees it as dead.
#![allow(dead_code)]
#![allow(clippy::cargo_common_metadata)]

use std::collections::HashMap;

struct Solution;

fn main() {
    let _ = Solution {};
    Solution::two_sum(vec![2, 7, 11, 15], 9);
}

impl Solution {
    // `nums` is taken by value to match LeetCode's published method signature.
    #[allow(clippy::needless_pass_by_value)]
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut seen = HashMap::new();
        for (i, num) in nums.iter().enumerate() {
            let need = target - num;
            if let Some(&j) = seen.get(&need) {
                return vec![
                    i32::try_from(j).expect("index fits in i32"),
                    i32::try_from(i).expect("index fits in i32"),
                ];
            }
            seen.insert(num, i);
        }

        vec![0, 0]
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(&[2, 7, 11, 15], 9, &[0, 1])]
    #[case(&[3, 2, 4], 6, &[1, 2])]
    #[case(&[3, 3], 6, &[0, 1])]
    #[case(&[-1, -2, -3, -4, -5], -8, &[2, 4])]
    #[case(&[3, 2, 95, 4, -3], 92, &[2, 4])]
    fn test_solution(#[case] nums: &[i32], #[case] target: i32, #[case] expected: &[i32]) {
        assert_eq!(Solution::two_sum(nums.to_vec(), target), expected);
    }
}
