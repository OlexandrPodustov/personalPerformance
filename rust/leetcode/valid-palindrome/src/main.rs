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

        vc.iter()
            .filter(|&s| s.is_alphanumeric())
            .eq(vc.iter().filter(|&s| s.is_alphanumeric()).rev())
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
