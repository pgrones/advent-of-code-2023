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

    let mut sorted_maps = Vec::new();

    for mut map in maps {
        map.sort_by(|a, b| a.src_range_start.cmp(&b.src_range_start));
        sorted_maps.push(map);
    }

    let mut result = part1(seeds.clone(), sorted_maps.clone());
    println!("Part 1: {}", result);

    result = part2_chunks(seeds, sorted_maps);
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

fn part2_chunks(seeds: Vec<i64>, maps: Vec<Vec<Instruction>>) -> i64 {
    let mut i = 0;
    let mut chunks = Vec::new();

    while i < seeds.len() {
        let min = seeds[i];
        let max = seeds[i] + seeds[i + 1];

        chunks.push((min, max));

        i += 2;
    }

    for map in &maps {
        let mut mapped_chunks = Vec::new();

        for (a, b) in &chunks {
            let mut start_index = *a;

            while start_index != *b {
                let mut new_chunk;

                let instruction = map.iter().find(|x| {
                    start_index >= x.src_range_start && start_index < (x.src_range_start + x.range)
                });

                if instruction.is_some() {
                    let inst = instruction.unwrap();

                    new_chunk = (start_index, cmp::min(inst.src_range_start + inst.range, *b));

                    let offset = inst.dest_range_start - inst.src_range_start;

                    start_index = new_chunk.1;

                    new_chunk.0 += offset;
                    new_chunk.1 += offset;
                } else {
                    let instruction = map.iter().find(|x| x.src_range_start > start_index);

                    if instruction.is_some() {
                        let inst = instruction.unwrap();

                        new_chunk = (start_index, cmp::min(inst.src_range_start, *b));
                    } else {
                        new_chunk = (start_index, *b);
                    }

                    start_index = new_chunk.1;
                }

                mapped_chunks.push(new_chunk);
            }
        }

        chunks = mapped_chunks;
    }

    chunks.sort_by(|a, b| a.0.cmp(&b.0));

    chunks[0].0
}
