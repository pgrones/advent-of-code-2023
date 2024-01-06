#![allow(dead_code, unused_imports)]

use std::{
    fs::File,
    io::{self, BufReader, Lines},
    thread, time,
};

use crate::utils::read_lines;

const STACK_SIZE: usize = 4 * 1024 * 1024;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day23/input_{run_as}.txt");

    let lines = read_lines(&input_file);
    let hiking_map = lines
        .iter()
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    // We need a bigger stack for this
    let child = thread::Builder::new()
        .stack_size(STACK_SIZE)
        .spawn(move || println!("Part 1: {}", part1(&hiking_map)))
        .unwrap();

    // Wait for thread to join
    child.join().unwrap();

    // count = part2(lines2);
    // println!("Part 2: {}", count);

    Ok(())
}

fn part1(hiking_map: &Vec<Vec<char>>) -> u32 {
    let start = (1usize, 1usize);
    let last_pos = (1usize, 0usize);
    let end = (hiking_map[0].len() - 2, hiking_map.len() - 1);

    let mut steps = vec![0];

    traverse(hiking_map, last_pos, start, end, &mut steps, 0);

    *steps.iter().max().unwrap()
}

fn traverse(
    hiking_map: &Vec<Vec<char>>,
    last_pos: (usize, usize),
    curr_pos: (usize, usize),
    end_pos: (usize, usize),
    steps: &mut Vec<u32>,
    step_pos: usize,
) {
    steps[step_pos] += 1;

    if curr_pos == end_pos {
        return;
    }

    // render_map(hiking_map, curr_pos, steps);
    // thread::sleep(time::Duration::from_millis(200));

    let paths = get_next_paths(hiking_map, last_pos, curr_pos);

    let steps_copy = steps.clone();

    for i in 0..paths.len() {
        let mut pos = step_pos;

        if i > 0 {
            steps.push(steps_copy[step_pos]);
            pos = steps.len() - 1;
        }

        traverse(hiking_map, curr_pos, paths[i], end_pos, steps, pos);
    }
}

fn get_next_paths(
    hiking_map: &Vec<Vec<char>>,
    last_pos: (usize, usize),
    curr_pos: (usize, usize),
) -> Vec<(usize, usize)> {
    let mut result = vec![];

    if curr_pos.1 > 0
        && (curr_pos.0, curr_pos.1 - 1) != last_pos
        && hiking_map[curr_pos.1 - 1][curr_pos.0] != '#'
        && hiking_map[curr_pos.1 - 1][curr_pos.0] != 'v'
    {
        result.push((curr_pos.0, curr_pos.1 - 1));
    }

    if curr_pos.1 < hiking_map.len() - 1
        && (curr_pos.0, curr_pos.1 + 1) != last_pos
        && hiking_map[curr_pos.1 + 1][curr_pos.0] != '#'
        && hiking_map[curr_pos.1 + 1][curr_pos.0] != '^'
    {
        result.push((curr_pos.0, curr_pos.1 + 1));
    }

    if curr_pos.0 > 0
        && (curr_pos.0 - 1, curr_pos.1) != last_pos
        && hiking_map[curr_pos.1][curr_pos.0 - 1] != '#'
        && hiking_map[curr_pos.1][curr_pos.0 - 1] != '>'
    {
        result.push((curr_pos.0 - 1, curr_pos.1));
    }

    if curr_pos.0 < hiking_map[0].len() - 1
        && (curr_pos.0 + 1, curr_pos.1) != last_pos
        && hiking_map[curr_pos.1][curr_pos.0 + 1] != '#'
        && hiking_map[curr_pos.1][curr_pos.0 + 1] != '<'
    {
        result.push((curr_pos.0 + 1, curr_pos.1));
    }

    result
}

fn render_map(hiking_map: &Vec<Vec<char>>, curr_pos: (usize, usize), steps: &Vec<u32>) {
    for y in 0..hiking_map.len() {
        for x in 0..hiking_map[0].len() {
            if curr_pos == (x, y) {
                print!("O");
            } else {
                print!("{}", hiking_map[y][x]);
            }
        }
        println!();
    }
    println!("{:?}", steps);
}

fn part2(lines: Lines<BufReader<File>>) -> u32 {
    0
}
