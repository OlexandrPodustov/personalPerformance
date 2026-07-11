pub struct Solution;

impl Solution {
    pub fn merge(nums1: &mut [i32], m: i32, nums2: &mut [i32], n: i32) {
        let mut n1 = m - 1;
        let mut n2 = n - 1;
        let mut write_at = nums1.len() - 1;
        while n2 >= 0 {
            if n1 >= 0 && nums1[n1 as usize] > nums2[n2 as usize] {
                nums1[write_at] = nums1[n1 as usize];
                n1 -= 1;
            } else {
                nums1[write_at] = nums2[n2 as usize];
                n2 -= 1;
            }

            write_at = write_at.saturating_sub(1);
        }
    }
}
fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(vec![1, 2, 3, 0, 0, 0], 3, vec![2, 5, 6], 3, vec![1, 2, 2, 3, 5, 6])]
    #[case(vec![1], 1, vec![], 0, vec![1])]
    #[case(vec![2, 0], 1, vec![1], 1, vec![1, 2])]
    fn it_works(
        #[case] mut nums1: Vec<i32>,
        #[case] m: i32,
        #[case] mut nums2: Vec<i32>,
        #[case] n: i32,
        #[case] expected_result: Vec<i32>,
    ) {
        Solution::merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, expected_result);
    }
}
