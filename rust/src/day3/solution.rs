use std::io;

use regex::Regex;

use crate::utils::read_lines;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day3/input_{run_as}.txt");

    let lines = read_lines(&input_file);

    let mut result = part1(lines.clone());
    println!("Part 1: {}", result);

    result = part2(lines);
    println!("Part 2: {}", result);

    Ok(())
}

fn part1(lines: Vec<String>) -> u32 {
    let mut result = 0;

    // loop over all lines
    for line_index in 0..lines.len() {
        // get the coordinates of all numbers on a line
        let numbers: Vec<(usize, usize)> = Regex::new(r"\d+")
            .unwrap()
            .find_iter(&lines[line_index])
            .map(|x| (x.start(), x.end()))
            .collect();

        for (start_index, end_index) in numbers {
            // create min/max indexes to check
            let min_y;
            if line_index == 0 {
                min_y = 0;
            } else {
                min_y = line_index - 1;
            }

            let max_y;
            if line_index == lines.len() - 1 {
                max_y = lines.len();
            } else {
                max_y = line_index + 2;
            }

            let min_x;
            if start_index == 0 {
                min_x = start_index;
            } else {
                min_x = start_index - 1;
            }

            let max_x;
            if end_index == lines[line_index].len() {
                max_x = end_index;
            } else {
                max_x = end_index + 1;
            }

            // loop over all places to check
            'outer: for y in min_y..max_y {
                for x in min_x..max_x {
                    // skip the number itself
                    if y == line_index && x >= start_index && x < end_index {
                        continue;
                    }

                    // if there is a symbol (no ".") the number is valid
                    if lines[y].chars().nth(x).unwrap() != '.' {
                        match lines[line_index][start_index..end_index].parse::<u32>() {
                            Err(err) => println!("{}", err),
                            Ok(v) => result += v,
                        }

                        break 'outer;
                    }
                }
            }
        }
    }

    result
}

fn part2(lines: Vec<String>) -> u32 {
    let mut result = 0;

    // loop over all lines
    for line_index in 0..lines.len() {
        // get the coordinates of all * on a line
        let gears: Vec<usize> = Regex::new(r"\*")
            .unwrap()
            .find_iter(&lines[line_index])
            .map(|x| x.start())
            .collect();

        for index in gears {
            // create min/max indexes to check
            let min_y;
            if line_index == 0 {
                min_y = 0;
            } else {
                min_y = line_index - 1;
            }

            let max_y;
            if line_index == lines.len() - 1 {
                max_y = lines.len();
            } else {
                max_y = line_index + 2;
            }

            let min_x;
            if index == 0 {
                min_x = index;
            } else {
                min_x = index - 1;
            }

            let max_x;
            if index == lines[line_index].len() {
                max_x = index;
            } else {
                max_x = index + 2;
            }

            // keep track of the indexes and the numbers

            let mut numbers: Vec<u32> = Vec::new();

            // loop over all places to check
            'outer: for y in min_y..max_y {
                let mut checked_indexes: Vec<usize> = Vec::new();

                for x in min_x..max_x {
                    // skip the gear itself and digits that were part of a previous number
                    if (y == line_index && x == index) || checked_indexes.contains(&x) {
                        continue;
                    }

                    // if there is a digit
                    if lines[y].chars().nth(x).unwrap().is_digit(10) {
                        // get the start and end indexes of the number
                        let mut start = x;
                        while start > 0 && lines[y].chars().nth(start - 1).unwrap().is_digit(10) {
                            start -= 1;
                            checked_indexes.push(start);
                        }

                        let mut end = x;
                        while end < lines[y].len() - 1
                            && lines[y].chars().nth(end + 1).unwrap().is_digit(10)
                        {
                            end += 1;
                            checked_indexes.push(end);
                        }

                        checked_indexes.push(x);

                        // add the number to a list to add them later
                        match lines[y][start..end + 1].parse::<u32>() {
                            Err(err) => println!("{}", err),
                            Ok(v) => numbers.push(v),
                        }

                        // only gears with exactly 2 numbers are valid, so we can return early if there are more than 2
                        if numbers.len() == 3 {
                            break 'outer;
                        }
                    }
                }
            }

            if numbers.len() == 2 {
                result += numbers[0] * numbers[1];
            }
        }
    }

    result
}
