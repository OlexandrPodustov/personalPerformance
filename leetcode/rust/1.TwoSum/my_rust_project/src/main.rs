struct Solution; // This defines the empty struct

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut f: i32 = -1;
        let mut s: i32 = -1;
        'outer: for (i, num) in nums.iter().enumerate() {
            let diff = target - num;
            // println!("i {} - nums.len()  {} diff {}", i, nums.len(), diff);
            if i + 1 < nums.len() {
                for (j, numb) in nums[i + 1..].iter().enumerate() {
                    // println!("j, numb {} - {}", j, numb);

                    if diff == *numb {
                        f = i as i32;
                        s = j as i32;
                        s += f + 1;
                        break 'outer;
                    }
                }
            }
        }

        vec![f, s]
    }
}

fn main() {
    run_solution(vec![2, 7, 11, 15], 9);
    run_solution(vec![3, 2, 4], 6);
    run_solution(vec![3, 3], 6);
}

// Vec<i32>, target: i32
fn run_solution(numbers: Vec<i32>, target_val: i32) {
    println!(
        "Calling two_sum with numbers: {:?} and target: {}",
        numbers, target_val
    );

    let result = Solution::two_sum(numbers, target_val);
    println!("Result: {:?}", result);
}
