use std::{
    fs::File,
    io::{self, BufReader, Lines},
};

use crate::utils::read_lines_iterable;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day9/input_{run_as}.txt");

    let lines = read_lines_iterable(input_file.clone())?;

    let count = part1_and_2(lines);
    println!("Part 1: {}", count.0);
    println!("Part 2: {}", count.1);

    Ok(())
}

pub fn part1_and_2(lines: Lines<BufReader<File>>) -> (i32, i32) {
    let mut result = (0, 0);

    for line in lines {
        if let Ok(l) = line {
            // get the numbers from the line and put them into a 2D array
            // an array of sequences that converge towards all zeros
            let mut sequences: Vec<Vec<i32>> =
                vec![l.split(" ").map(|x| x.parse::<i32>().unwrap()).collect()];

            // calculate the difference between each number pair
            while sequences.last().unwrap().iter().any(|&x| x != 0) {
                let last_seq = sequences.last().unwrap();

                let mut next_seq = vec![];
                for i in 0..last_seq.len() - 1 {
                    next_seq.push(last_seq[i + 1] - last_seq[i]);
                }

                sequences.push(next_seq);
            }

            // part 1: sum up the last values
            result.0 += sequences.iter().fold(0, |acc, x| acc + x[x.len() - 1]);

            // part2: subtract the first values
            result.1 += sequences.iter().rev().fold(0, |acc, x| x[0] - acc);
        }
    }

    result
}
