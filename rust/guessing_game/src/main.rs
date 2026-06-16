use colored::Colorize;
use std::cmp::Ordering;
use std::io;

fn main() {
    let sec = rand::random_range(1..=101);

    println!("Guess the number!");
    println!("Please input your guess.");

    loop {
        let mut guess = String::new();
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(err) => {
                println!("Please input a number! {err}");
                continue;
            }
        };

        println!("You guessed: {guess}");

        match guess.cmp(&sec) {
            Ordering::Less => println!("{}", "Too small!".red()),
            Ordering::Greater => println!("{}", "Too big!".red()),
            Ordering::Equal => {
                println!("{}", "You win!".green());
                break;
            }
        }
    }
}
