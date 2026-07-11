pub struct Solution;

impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut read_index: usize = 1;
        let mut write_index: usize = 1;
        while read_index < nums.len() {
            if nums[read_index - 1] == nums[read_index] {
                read_index += 1;
                continue;
            } else {
                nums[write_index] = nums[read_index];
                read_index += 1;
                write_index += 1;
            }
        }

        write_index as i32
    }
}
#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(vec![1, 1, 2], 2)]
    #[case(vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4], 5)]
    fn it_works(#[case] mut nums: Vec<i32>, #[case] k: i32) {
        let res = Solution::remove_duplicates(&mut nums);
        assert_eq!(res, k);
    }
}
