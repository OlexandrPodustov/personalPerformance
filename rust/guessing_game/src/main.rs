use rand::Rng;
use std::io;
fn main() {
    println!("Guess the number!");

    let sec = rand::thread_rng().gen_range(1..=100);

    println!("Please input your guess.");

    let mut guess = String::new();
    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");
    println!("You guessed: {}", guess);

    if guess.trim() == sec.to_string() {
        println!("You guessed it right!");
        return;
    }
    println!("You failed to guess it");
}
