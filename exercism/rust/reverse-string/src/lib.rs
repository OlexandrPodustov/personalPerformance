use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    if input.is_empty() {
        return String::new();
    }

    input.graphemes(true).rev().collect()
}
