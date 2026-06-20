#![allow(clippy::cargo_common_metadata)]

pub struct Solution;

impl Solution {
    #[must_use]
    pub fn is_palindrome(s: String) -> bool {
        let vc: Vec<char> = s
            .chars()
            .filter(|&s| s.is_alphanumeric())
            .map(|c| c.to_ascii_lowercase())
            .collect();

        if vc.is_empty() {
            return true;
        }

        let end = vc.len();
        let (mut s, mut e) = (0, end - 1);
        loop {
            if s >= e {
                break;
            }

            if vc[s] != vc[e] {
                return false;
            }
            s += 1;
            e -= 1;
        }

        true
    }
}
fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::is_palindrome("A man, a plan, a canal: Panama".to_owned()),
            true
        );
        assert_eq!(Solution::is_palindrome("race a car".to_owned()), false);
        assert_eq!(Solution::is_palindrome(" ".to_owned()), true);
        assert_eq!(Solution::is_palindrome("ab cd".to_owned()), false);
        assert_eq!(Solution::is_palindrome("ab cd".to_owned()), false);
    }
}
