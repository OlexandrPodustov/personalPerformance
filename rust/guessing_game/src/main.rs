use rand::Rng;
use std::io;
fn main() {
    println!("Guess the number!");
    println!("Please input your guess.");

    let mut guess = String::new();
    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");
    println!("You guessed: {}", guess);

    let mut rnd = rand::thread_rng();
    let sec = rnd.gen_range(1..=100);

    if guess == sec.to_string() {
        println!("You guessed it right!");
    }
    println!("You failed to guess it - {}", sec);
}
