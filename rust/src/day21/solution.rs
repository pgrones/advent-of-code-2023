use std::{
    collections::HashSet,
    fs::File,
    io::{self, BufReader, Lines},
    vec,
};

use crate::utils::read_lines;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day21/input_{run_as}.txt");

    let lines = read_lines(&input_file);
    let garden = lines
        .iter()
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut count = part1(&garden);
    println!("Part 1: {}", count);

    // let lines2 = read_lines_iterable(input_file)?;
    // count = part2(lines2);
    // println!("Part 2: {}", count);

    Ok(())
}

fn part1(garden: &Vec<Vec<char>>) -> u32 {
    let mut start = (0, 0);

    'outer: for y in 0..garden.len() {
        for x in 0..garden[0].len() {
            if garden[y][x] == 'S' {
                start = (x, y);
                break 'outer;
            }
        }
    }

    let mut queue = vec![start];

    for _ in 0..64 {
        let mut next_positions = HashSet::new();

        while queue.len() > 0 {
            next_positions.extend(execute_step(garden, queue.pop().unwrap()));
        }

        queue = Vec::from_iter(next_positions)
    }

    return queue.len() as u32;
}

fn execute_step(garden: &Vec<Vec<char>>, curr_pos: (usize, usize)) -> Vec<(usize, usize)> {
    let mut result = vec![];

    if curr_pos.1 > 0 && garden[curr_pos.1 - 1][curr_pos.0] != '#' {
        result.push((curr_pos.0, curr_pos.1 - 1));
    }

    if curr_pos.1 < garden.len() - 1 && garden[curr_pos.1 + 1][curr_pos.0] != '#' {
        result.push((curr_pos.0, curr_pos.1 + 1));
    }

    if curr_pos.0 > 0 && garden[curr_pos.1][curr_pos.0 - 1] != '#' {
        result.push((curr_pos.0 - 1, curr_pos.1));
    }

    if curr_pos.0 < garden[0].len() - 1 && garden[curr_pos.1][curr_pos.0 + 1] != '#' {
        result.push((curr_pos.0 + 1, curr_pos.1));
    }

    result
}

// fn render_garden(garden: &Vec<Vec<char>>, plots: &Vec<(usize, usize)>) {
//     for y in 0..garden.len() {
//         for x in 0..garden[0].len() {
//             if plots.contains(&(x, y)) {
//                 print!("O");
//             } else {
//                 print!("{}", garden[y][x]);
//             }
//         }
//         println!();
//     }
//     println!();
// }

fn part2(lines: Lines<BufReader<File>>) -> u32 {
    0
}
