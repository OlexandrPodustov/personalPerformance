#![allow(clippy::cargo_common_metadata)]

// fn main() {
//     let spacey_num = "  42  ";

//     // Your code here (shadow 'spacey_num' until it is an i32)

//     let spacey_num = 84;

//     println!("The doubled value is: {}", spacey_num); // Should print 84
// }

// use std::thread;

// fn main() {
//     let message = String::from("Secret Data");

//     let message2 = message.clone();
//     let handle = thread::spawn(move || {
//         // Task: Print the message here
//         println!("1 {}", message2);
//     });

//     println!("2 {}", message);

//     // Task: Try to print 'message' here and observe the compiler error.
//     // Then, join the thread handle to ensure it finishes.
//     handle.join().unwrap();
// }

// use std::sync::{Arc, Mutex};
// use std::thread;

// fn main() {
//     let counter = Arc::new(Mutex::new(0));
//     let mut handles = vec![];

//     for _ in 0..10 {
//         let counter = Arc::clone(&counter);
//         let handle = thread::spawn(move || {
//             // Task: Lock the mutex and increment the inner value
//             let mut count = counter.lock().unwrap();
//             *count += 1;
//         });
//         handles.push(handle);
//     }

//     for handle in handles {
//         handle.join().unwrap();
//     }

//     println!("Result: {}", *counter.lock().unwrap());
// }

/// Returns the largest integer in `numbers`, or `None` if the slice is empty.
fn largest(numbers: &[i32]) -> Option<i32> {
    numbers.iter().copied().max()
}

fn main() {
    let nums = vec![3, 7, 2, 9, 1];
    match largest(&nums) {
        Some(n) => println!("Largest: {n}"),
        None => println!("Empty list"),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn largest_nonempty() {
        assert_eq!(largest(&[3, 7, 2, 9, 1]), Some(9));
        assert_eq!(largest(&[42]), Some(42));
        assert_eq!(largest(&[-5, -1, -10]), Some(-1));
    }

    #[test]
    fn largest_empty() {
        assert_eq!(largest(&[]), None);
    }
}
