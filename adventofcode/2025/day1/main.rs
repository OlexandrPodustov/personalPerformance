#![allow(clippy::cargo_common_metadata)]

use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let mut result = 0;
    let mut points_to = 50;

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let input_line = line?;
        let (direction, distance) = input_line.split_at(1);
        let distance_int = distance.parse::<i32>().unwrap();
        let distance_mod = distance_int % 100;

        println!(
            "direction: {direction}, distance: {distance_int}, distance modulo: {distance_mod}"
        );

        if direction == "L" {
            points_to -= distance_mod;
        } else {
            points_to += distance_mod;
        }
        if points_to < 0 {
            points_to += 100;
        } else if points_to >= 100 {
            points_to -= 100;
        }
        println!("  points_to: {points_to}");

        if points_to == 0 {
            result += 1;
        }
    }

    println!("result: {result}");

    Ok(())
}
