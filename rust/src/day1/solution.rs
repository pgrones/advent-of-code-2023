use std::{
    fs::File,
    io::{self, BufReader, Lines},
};

use crate::utils::read_lines_iterable;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day1/input_{run_as}.txt");

    let lines = read_lines_iterable(input_file.clone())?;

    let mut count = part1(lines);
    println!("Part 1: {}", count);

    let lines2 = read_lines_iterable(input_file)?;
    count = part2(lines2);
    println!("Part 2: {}", count);

    Ok(())
}

pub fn part1(lines: Lines<BufReader<File>>) -> u32 {
    let mut count: u32 = 0;

    for line in lines {
        if let Ok(l) = line {
            let mut idx1: usize = 0;
            let mut idx2: usize = l.len() - 1;

            let mut c1: char;
            let mut c2: char;

            let number1: u32;
            let number2: u32;

            let char_list: Vec<char> = l.chars().collect::<Vec<char>>();

            loop {
                c1 = char_list[idx1];
                let n1 = c1.to_digit(10);

                if n1.is_some() {
                    number1 = n1.unwrap();
                    break;
                }

                idx1 += 1;
            }

            loop {
                c2 = char_list[idx2];
                let n2 = c2.to_digit(10);

                if n2.is_some() {
                    number2 = n2.unwrap();
                    break;
                }
                idx2 -= 1;
            }

            count += 10 * number1 + number2;
        }
    }

    count
}

pub fn part2(lines: Lines<BufReader<File>>) -> u32 {
    let mut count: u32 = 0;

    let spelled_digits = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    for line in lines {
        if let Ok(l) = line {
            let mut idx1: usize = 0;
            let mut idx2: usize = l.len() - 1;

            let mut c1: char;
            let mut c2: char;

            let number1: u32;
            let number2: u32;

            let char_list: Vec<char> = l.chars().collect::<Vec<char>>();

            'outer: loop {
                c1 = char_list[idx1];
                let n1 = c1.to_digit(10);

                if n1.is_some() {
                    number1 = n1.unwrap();
                    break;
                }

                for (i, digit) in spelled_digits.iter().enumerate() {
                    let s: String = char_list[idx1..].into_iter().collect();
                    if s.starts_with(digit) {
                        number1 = i as u32 + 1;
                        break 'outer;
                    }
                }

                idx1 += 1;
            }

            'outer: loop {
                c2 = char_list[idx2];
                let n2 = c2.to_digit(10);

                if n2.is_some() {
                    number2 = n2.unwrap();
                    break;
                }

                for (i, digit) in spelled_digits.iter().enumerate() {
                    let s: String = char_list[idx2..].into_iter().collect();
                    if s.starts_with(digit) {
                        number2 = i as u32 + 1;
                        break 'outer;
                    }
                }

                idx2 -= 1;
            }

            count += 10 * number1 + number2;
        }
    }

    count
}
