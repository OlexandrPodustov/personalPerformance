use std::io::{self, BufRead};

struct Point {
    points_to: i32,
}

impl Point {
    fn new() -> Self {
        Point { points_to: 50 }
    }

    fn rotate(&mut self, direction: &str, distance: i32) -> i32 {
        let mut pointed_to_zero: i32 = 0;

        println!("before points_to: {}", self.points_to);

        let distance_mod = distance % 100;

        if direction == "L" {
            pointed_to_zero += (self.points_to - distance) / 100;
            if (self.points_to - distance) == 0 {
                pointed_to_zero += 1;
            }

            println!(
                "pointed_to_zero {} while rotating left from: {} - {} = {}",
                pointed_to_zero,
                self.points_to,
                distance,
                self.points_to - distance,
            );
            self.points_to -= distance_mod;
        } else {
            pointed_to_zero += (self.points_to + distance) / 100;
            println!(
                "pointed_to_zero {} while rotating right from: {} + {} = {}",
                pointed_to_zero,
                self.points_to,
                distance,
                self.points_to + distance,
            );

            self.points_to += distance_mod;
        }
        if self.points_to < 0 {
            self.points_to += 100;
        } else if self.points_to >= 100 {
            self.points_to -= 100;
        }

        // if self.points_to - distance < 0 {
        //     pointed_to_zero += 1;
        // }

        println!("pointed_to_zero times: {}", pointed_to_zero);
        println!("after points_to: {}", self.points_to);

        pointed_to_zero
    }
}

fn main() -> io::Result<()> {
    let mut result = 0;
    let mut current_point: Point = Point::new();

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let input_line = line?;
        if input_line.trim().is_empty() {
            // break the loop on empty line
            break;
        }

        let (direction, distance) = input_line.split_at(1);
        let distance_int = distance.parse::<i32>().unwrap();

        println!("direction: {}, distance: {}", direction, distance_int);

        let zero_times = current_point.rotate(direction, distance_int);
        println!("zero_times: {}", zero_times);

        result += zero_times;
        println!("result accumulated: {}", result);
        println!(" ------------------- ");
    }

    println!("result: {}", result);

    Ok(())
}
