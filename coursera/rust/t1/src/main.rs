struct Highlight<'a> {
    content: &'a str,
}

fn main() {
    let text = String::from("The quick brown fox");
    let fox = Highlight { content: &text[10..15] };
    println!("Highlighted word: {}", fox.content);
}

