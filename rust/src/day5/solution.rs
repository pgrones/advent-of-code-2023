use regex::Regex;
use std::{cmp, i64::MAX, io};

use crate::utils::{read_lines, read_lines_into_lists_of_structs, Instantiable};

#[derive(Debug, Clone)]
struct Instruction {
    src_range_start: i64,
    dest_range_start: i64,
    range: i64,
}

impl Instantiable for Instruction {
    fn new(values: Vec<&str>) -> Self {
        Instruction {
            dest_range_start: values[0].parse::<i64>().unwrap(),
            src_range_start: values[1].parse::<i64>().unwrap(),
            range: values[2].parse::<i64>().unwrap(),
        }
    }
}

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day5/input_{run_as}.txt");

    let seeds: Vec<i64> = read_lines(&input_file)[0]
        .strip_prefix("seeds: ")
        .unwrap()
        .split(" ")
        .map(|x| x.parse::<i64>().unwrap())
        .collect();

    let maps = read_lines_into_lists_of_structs::<Instruction, _>(
        &input_file,
        Regex::new(r"\s").unwrap(),
        "",
        |x, i| i > 1 && !x.ends_with(":"),
    );

    let mut result = part1(seeds.clone(), maps.clone());
    println!("Part 1: {}", result);

    result = part2(seeds, maps);
    println!("Part 2: {}", result);

    Ok(())
}

fn part1(seeds: Vec<i64>, maps: Vec<Vec<Instruction>>) -> i64 {
    let mut result = MAX;

    for seed in seeds {
        let mut value = seed;

        for map in &maps {
            let instruction = map
                .iter()
                .find(|x| value >= x.src_range_start && value < x.src_range_start + x.range);

            if instruction.is_none() {
                continue;
            }

            let offset =
                instruction.unwrap().dest_range_start - instruction.unwrap().src_range_start;

            value += offset;
        }

        if value < result {
            result = value;
        }
    }

    result
}

fn part2(seeds: Vec<i64>, maps: Vec<Vec<Instruction>>) -> i64 {
    let mut result = MAX;
    let mut i = 0;

    while i < seeds.len() {
        let min = cmp::min(seeds[i], seeds[i + 1]);
        let max = cmp::max(seeds[i], seeds[i + 1]);

        for seed in min..max {
            let mut value = seed;

            for map in &maps {
                let instruction = map
                    .iter()
                    .find(|x| value >= x.src_range_start && value < x.src_range_start + x.range);

                if instruction.is_none() {
                    continue;
                }

                let offset =
                    instruction.unwrap().dest_range_start - instruction.unwrap().src_range_start;

                value += offset;
            }

            if value < result {
                result = value;
            }
        }

        i += 2;
    }

    result
}
