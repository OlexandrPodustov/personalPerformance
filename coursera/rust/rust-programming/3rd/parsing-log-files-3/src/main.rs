#![allow(clippy::cargo_common_metadata)]

use flate2::read::GzDecoder;
use regex::Regex;
use std::env;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::path::Path;

fn read_buffer(file_path: &str) {
    // Initialize variables for error rate calculation.
    let mut total_errors = 0;
    let mut hour_total = 0;
    let mut hour_errors = 0;
    let mut current_hour: Option<String> = None;

    // Syslog regex: captures the hour part of the timestamp and the message.
    // Standard format: "Mmm dd hh:mm:ss hostname message..."
    // Group 1: "Mmm dd hh" (e.g., "May  7 10")
    // Group 2: The rest of the line (message)
    let syslog_regex = Regex::new(r"^([A-Z][a-z]{2}\s+\d+\s+\d{2}):\d{2}:\d{2}\s+(.*)$").unwrap();

    let file = match File::open(file_path) {
        Ok(f) => f,
        Err(e) => {
            eprintln!("Error opening file {file_path}: {e}");
            return;
        }
    };

    let reader: Box<dyn BufRead> = if Path::new(file_path)
        .extension()
        .is_some_and(|ext| ext.eq_ignore_ascii_case("gz"))
    {
        let decompressor = GzDecoder::new(file);
        Box::new(BufReader::new(decompressor))
    } else {
        Box::new(BufReader::new(file))
    };

    for line in reader.lines() {
        let line = match line {
            Ok(line) => line,
            Err(error) => {
                eprintln!("Error reading line: {error}");
                continue;
            }
        };

        // Extract timestamp (hour resolution) and message from the syslog line.
        if let Some(captures) = syslog_regex.captures(&line) {
            let timestamp_hour = &captures[1];
            let message = &captures[2];

            // Check if the hour has changed.
            if current_hour.as_deref() != Some(timestamp_hour) {
                // Calculate and print error rate for the previous hour.
                if let Some(prev_hour) = current_hour {
                    let rate = if hour_total > 0 {
                        (f64::from(hour_errors) / f64::from(hour_total)) * 100.0
                    } else {
                        0.0
                    };
                    println!(
                        "{prev_hour} - Errors: {hour_errors}, Total: {hour_total}, Rate: {rate:.2}%"
                    );
                }

                // Reset counters for the new hour.
                hour_errors = 0;
                hour_total = 0;
                current_hour = Some(timestamp_hour.to_string());
            }

            hour_total += 1;
            // Check if the message contains "error" (case-insensitive) to identify errors.
            if message.to_lowercase().contains("error") {
                hour_errors += 1;
                total_errors += 1;
            }
        }
    }

    // Print error rate for the final hour in the file.
    if let Some(prev_hour) = current_hour {
        let rate = if hour_total > 0 {
            (f64::from(hour_errors) / f64::from(hour_total)) * 100.0
        } else {
            0.0
        };
        println!("{prev_hour} - Errors: {hour_errors}, Total: {hour_total}, Rate: {rate:.2}%");
    }

    println!("Total error count for current log: {total_errors}");
}
fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: log_error_rate <log_file_path>");
        std::process::exit(1);
    }

    let log_file_path = &args[1];

    read_buffer(log_file_path);
}
