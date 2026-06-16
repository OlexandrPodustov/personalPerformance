#![allow(clippy::cargo_common_metadata)]

use std::env;
use std::fs;
use std::path::Path;

fn walk_path(path: &Path) {
    if let Ok(entries) = fs::read_dir(path) {
        for entry in entries.flatten() {
            let path = entry.path();
            if path.is_dir() {
                println!("Dir: {}", path.display());
                walk_path(&path);
            } else {
                println!("File: {}", path.display());
            }
        }
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Usage: {} <path>", args[0]);
        return;
    }

    let path = Path::new(&args[1]);

    if !path.exists() {
        println!("Path '{}' does not exist.", path.display());
        if args[1] == "target" {
            println!(
                "Note: In a Cargo workspace, the 'target' directory is usually at the workspace root."
            );
            println!("Try: cargo run -- ../../../target");
        }
        return;
    }

    walk_path(path);
}
