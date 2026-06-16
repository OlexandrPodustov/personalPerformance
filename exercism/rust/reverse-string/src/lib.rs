#![allow(clippy::cargo_common_metadata)]
use unicode_segmentation::UnicodeSegmentation;

#[must_use] 
pub fn reverse(input: &str) -> String {
    if input.is_empty() {
        return String::new();
    }

    input.graphemes(true).rev().collect()
}
