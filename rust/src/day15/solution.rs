use std::io;

use crate::utils::{char_to_ascii, read_lines};

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day15/input_{run_as}.txt");

    let init_sequence = read_lines(&input_file)[0]
        .split(",")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut count = part1(&init_sequence);
    println!("Part 1: {}", count);

    count = part2(&init_sequence);
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(init_sequence: &Vec<Vec<char>>) -> u32 {
    init_sequence
        .iter()
        .fold(0, |acc, x| acc + x.iter().fold(0, |acc, &c| hash(acc, c)))
}

fn part2(lines: &Vec<Vec<char>>) -> u32 {
    0
}

fn hash(mut curr_value: u32, c: char) -> u32 {
    // Determine the ASCII code for the current character of the string.
    let ascii_code = char_to_ascii(c) as u32;
    // Increase the current value by the ASCII code you just determined.
    curr_value += ascii_code;
    // Set the current value to itself multiplied by 17.
    curr_value *= 17;
    // Set the current value to the remainder of dividing itself by 256.
    curr_value %= 256;
    // The current value is the output of the HASH algorithm
    curr_value
}
