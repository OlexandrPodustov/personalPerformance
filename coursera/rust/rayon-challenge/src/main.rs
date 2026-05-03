// use rayon::prelude::*;
// use std::time::Instant;

// fn main2() {
//     rayon::ThreadPoolBuilder::new()
//         .num_threads(2)
//         .build_global()
//         .unwrap();

//     let data = vec![1; 50000];
//     let start = Instant::now();
//     let parallel_sum: i32 = data.par_iter().map(|x| x * x).sum();
//     let end = Instant::now();
//     let pt = end.duration_since(start);
//     println!("parallel sum time taken: \t {:03?}", pt);

//     let start = Instant::now();
//     let sequential_sum: i32 = data.iter().map(|x| x * x).sum();
//     let end = Instant::now();
//     let st = end.duration_since(start);
//     println!("sequential sum time taken: \t {:03?}", st);

//     println!("Parallel sum: {}", parallel_sum);
//     println!("Sequential sum: {}", sequential_sum);
// }

use rayon::prelude::*;
use std::collections::HashSet;
use std::sync::{Arc, Mutex};
// use std::thread;
use std::time::Instant;

fn main() {
    // Track unique thread IDs that Rayon uses
    let thread_ids: Arc<Mutex<HashSet<String>>> = Arc::new(Mutex::new(HashSet::new()));
    // rayon::ThreadPoolBuilder::new()
    //     .num_threads(16)
    //     .build_global()
    //     .unwrap();

    let data = vec![1i32; 500_000];

    // --- Parallel run ---
    // let ids_clone = Arc::clone(&thread_ids);
    let start = Instant::now();

    let parallel_sum: i32 = data
        .par_iter()
        .map(|x| {
            // Record which OS thread is doing this work
            // let id = format!("{:?}", thread::current().id());
            // ids_clone.lock().unwrap().insert(id);
            x * x
        })
        .sum();

    let parallel_time = start.elapsed();

    let unique_threads = thread_ids.lock().unwrap().clone();

    println!("=== Parallel ===");
    println!("  Time:           {parallel_time:?}");
    println!("  Sum:            {parallel_sum}");
    println!("  Threads used:   {}", unique_threads.len());
    println!("  Thread IDs:     {unique_threads:?}");

    // --- Sequential run ---
    let start = Instant::now();
    let sequential_sum: i32 = data.iter().map(|x| x * x).sum();
    let sequential_time = start.elapsed();

    println!("\n=== Sequential ===");
    println!("  Time:           {sequential_time:?}");
    println!("  Sum:            {sequential_sum}");
    println!("  Threads used:   1 (main thread only)");

    // --- Summary ---
    println!("\n=== Summary ===");
    println!("  Logical CPUs available: {}", rayon::current_num_threads());
    println!("  Rayon threads spawned:  {}", unique_threads.len());

    let ratio = sequential_time.as_nanos() as f64 / parallel_time.as_nanos() as f64;
    println!("  Speedup (seq/par):      {ratio:.2}x");

    if ratio < 1.0 {
        println!("  → Parallel was SLOWER (overhead > benefit at this data size)");
    } else {
        println!("  → Parallel was faster");
    }
}
