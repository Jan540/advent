use std::fs;

mod day3;

fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() < 2 {
        println!("Please provide a day number and the problem number as arguments");
        return;
    }

    let day = args[1]
        .parse::<i32>()
        .expect("Failed to parse day number :(");
    let problem = args[2]
        .parse::<i32>()
        .expect("Failed to parse problem number :(");

    let file = fs::read_to_string(format!("in/day{day}.in")).expect("Failed to read file :(");
    let lines: Vec<&str> = file.lines().collect();

    println!("Day {}, problem {}", day, problem);

    let result = match day {
        3 => match problem {
            1 => day3::problem1::solve(lines),
            2 => day3::problem2::solve(lines),
            _ => panic!("Invalid problem number :("),
        },
        _ => panic!("Invalid day number :("),
    };

    println!("{}", result);
}
