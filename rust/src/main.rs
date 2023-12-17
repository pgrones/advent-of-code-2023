mod utils;

mod day1;
use day1::solve as day1_solve;

mod day3;
use day3::solve as day3_solve;

mod day5;
use day5::solve as day5_solve;

mod day7;
use day7::solve as day7_solve;

mod day9;
use day9::solve as day9_solve;

mod day11;
use day11::solve as day11_solve;

mod naughty_even_days;
use naughty_even_days::day10::solve as day10_solve;

use std::env;

fn main() {
    let args = env::args().collect::<Vec<String>>();
    let day = args[1].parse::<i32>().unwrap();
    let run_as = args[2].chars().collect::<Vec<char>>()[0];

    match match day {
        1 => day1_solve(run_as),
        3 => day3_solve(run_as),
        5 => day5_solve(run_as),
        7 => day7_solve(run_as),
        9 => day9_solve(run_as),
        10 => day10_solve(run_as),
        11 => day11_solve(run_as),
        _ => Ok(()),
    } {
        Err(err) => println!("{}", err),
        _ => {}
    }
}
