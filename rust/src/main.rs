mod utils;

mod day1;
use day1::solve as day1_solve;

mod day3;
use day3::solve as day3_solve;

mod day5;
use day5::solve as day5_solve;

use std::env;

fn main() {
    let args = env::args().collect::<Vec<String>>();
    let day = args[1].chars().collect::<Vec<char>>()[0]
        .to_digit(10)
        .unwrap();
    let run_as = args[2].chars().collect::<Vec<char>>()[0];

    match match day {
        1 => day1_solve(run_as),
        3 => day3_solve(run_as),
        5 => day5_solve(run_as),
        _ => Ok(()),
    } {
        Err(err) => println!("{}", err),
        _ => {}
    }
}
