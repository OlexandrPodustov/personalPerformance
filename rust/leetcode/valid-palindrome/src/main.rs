#![allow(clippy::cargo_common_metadata)]

pub struct Solution;

impl Solution {
    #[must_use]
    pub fn is_palindrome(s: String) -> bool {
        s.chars()
            .filter(|&s| s.is_alphanumeric())
            .map(|c| c.to_ascii_lowercase())
            .eq(s
                .chars()
                .filter(|&s| s.is_alphanumeric())
                .map(|c| c.to_ascii_lowercase())
                .rev())
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
    #[case("A man, a plan, a canal: Panama", true)]
    #[case("race a car", false)]
    #[case(" ", true)]
    #[case("ab cd", false)]
    fn it_works(#[case] input: String, #[case] expected: bool) {
        assert_eq!(Solution::is_palindrome(input), expected);
    }
}
