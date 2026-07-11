#![allow(clippy::cargo_common_metadata)]

use std::error::Error;
use std::fs;

pub struct Config {
    pub query: String,
    pub file_path: String,
}

impl Config {
    /// # Errors
    ///
    /// Returns an error if fewer than 3 arguments are provided.
    pub fn build(args: &[String]) -> Result<Self, &'static str> {
        if args.len() < 3 {
            return Err("not enough arguments");
        }

        let query = args[1].clone();
        let file_path = args[2].clone();

        Ok(Self { query, file_path })
    }
}

/// # Errors
///
/// Returns an error if the configured file can't be read.
pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    // --snip--
    let contents = fs::read_to_string(config.file_path)?;
    println!("contents: {contents}");

    Ok(())
}

#[must_use]
pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    println!("{query}");

    let res = vec![];
    for line in contents.lines() {
        // do something with line

        println!("lll::: {line}");
    }
    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[ignore = "search() is an unfinished exercise"]
    fn one_result() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.";

        assert_eq!(vec!["safe, fast, productive."], search(query, contents));
    }
}
