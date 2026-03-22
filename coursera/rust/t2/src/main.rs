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

use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            // Task: Lock the mutex and increment the inner value
            let mut count = counter.lock().unwrap();
            *count += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}
