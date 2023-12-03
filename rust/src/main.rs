mod utils;

mod day1;
use day1::solve;

use std::env;

fn main() {
    let run_as = env::args().collect::<Vec<String>>()[1]
        .chars()
        .collect::<Vec<char>>()[0];

    match solve(run_as) {
        Err(err) => println!("{}", err),
        _ => {}
    }
}
