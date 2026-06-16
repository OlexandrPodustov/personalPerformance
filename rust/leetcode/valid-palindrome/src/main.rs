#![allow(clippy::cargo_common_metadata)]

pub struct Solution;

impl Solution {
    #[must_use] 
    pub fn is_palindrome(_s: String) -> bool {
        false
    }
}
fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {}
}
