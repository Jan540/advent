use std::fs;

fn main() {
    let file = fs::read_to_string("in/day1.in").expect("Failed to read file :(");

    for line in file.lines() {
        println!("{}", line)
    }
}
