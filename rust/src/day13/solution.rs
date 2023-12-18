use std::io::{self};

use crate::utils::read_lines_iterable;
#[derive(Debug)]
struct Pattern {
    rows: Vec<String>,
    cols: Vec<String>,
}

impl Pattern {
    fn new() -> Pattern {
        Pattern {
            rows: vec![],
            cols: vec![],
        }
    }
}

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day13/input_{run_as}.txt");

    let lines = read_lines_iterable(&input_file)?;

    let mut patterns = vec![];
    let mut pattern = Pattern::new();
    for line in lines {
        if let Ok(l) = line {
            if l == "" {
                patterns.push(pattern);
                pattern = Pattern::new();
                continue;
            }

            pattern.rows.push(l);
        }
    }

    patterns.push(pattern);

    // pivot the map to get the columns
    for p in &mut patterns {
        for x in 0..p.rows[0].len() {
            let mut col = "".to_string();

            for y in 0..p.rows.len() {
                col.push(p.rows[y].chars().collect::<Vec<char>>()[x]);
            }

            p.cols.push(col);
        }
    }

    let mut count = part1(&patterns);
    println!("Part 1: {}", count);

    count = part2(&patterns);
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(patterns: &Vec<Pattern>) -> i32 {
    let mut result = 0;

    for pattern in patterns {
        result += find_reflection(&pattern.cols, &pattern.rows);
    }

    result
}

fn part2(patterns: &Vec<Pattern>) -> i32 {
    let mut result = 0;

    for pattern in patterns {
        result += find_smudges(&pattern.cols, &pattern.rows);
    }

    result
}

fn find_reflection(lines: &Vec<String>, pivoted_lines: &Vec<String>) -> i32 {
    'outer: for i in 1..lines.len() {
        for j in i..lines.len() {
            // if everything until the left edge was valid, we end up here - that's our reflection line
            if 2 * i - j == 0 {
                return i as i32;
            }

            let left = &lines[2 * i - j - 1];
            let right = &lines[j];

            // if the pair doesn't match, got to the next pair
            if left != right {
                continue 'outer;
            }
        }

        // if everything until the right edge was valid, we end up here
        return i as i32;
    }

    // otherwise there is no reflection on this axis, so do the same for the other axis
    find_reflection(pivoted_lines, lines) * 100
}

fn find_smudges(lines: &Vec<String>, pivoted_lines: &Vec<String>) -> i32 {
    let mut smudges = 0;

    'outer: for i in 1..lines.len() {
        for j in i..lines.len() {
            // if everything until the left edge was valid, we end up here - that's our reflection line
            // as long as we encountered exactly one smudge
            if 2 * i - j == 0 {
                if smudges == 1 {
                    return i as i32;
                }

                continue 'outer;
            }

            let left = &lines[2 * i - j - 1];
            let right = &lines[j];

            // if the pair doesn't match, got to the next pair
            if left != right && smudges > 0 {
                smudges = 0;
                continue 'outer;
            } else if left != right && smudges == 0 {
                for i in 0..left.len() {
                    if smudges > 1 {
                        smudges = 0;
                        continue 'outer;
                    }

                    if left.chars().nth(i) != right.chars().nth(i) {
                        smudges += 1;
                    }
                }
            }
        }

        // if everything until the right edge was valid, we end up here
        // as long as we encountered exactly one smudge
        if smudges == 1 {
            return i as i32;
        }
    }

    // otherwise there is no reflection on this axis, so do the same for the other
    find_smudges(pivoted_lines, lines) * 100
}
