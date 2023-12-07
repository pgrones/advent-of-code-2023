use std::{
    fs::File,
    io::{self, BufReader, Lines},
};

use crate::utils::read_lines_iterable;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day0/input_{run_as}.txt");

    let lines = read_lines_iterable(input_file.clone())?;

    let mut count = part1(lines);
    println!("Part 1: {}", count);

    let lines2 = read_lines_iterable(input_file)?;
    count = part2(lines2);
    println!("Part 2: {}", count);

    Ok(())
}

pub fn part1(lines: Lines<BufReader<File>>) -> u32 {}

pub fn part2(lines: Lines<BufReader<File>>) -> u32 {}
