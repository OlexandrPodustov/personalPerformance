pub struct Solution;

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut cnt = 0;
        for i in 0..nums.len() {
            if nums[i] != val {
                nums[cnt] = nums[i];
                cnt += 1;
            }
        }

        cnt as i32
    }
}
const fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut nums = vec![3, 2, 2, 3];
        let val = 3;
        let result = Solution::remove_element(&mut nums, val);
        assert_eq!(result, 2);
    }
}
