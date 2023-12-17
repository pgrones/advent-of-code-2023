#![allow(dead_code)]

use std::{
    cmp,
    collections::HashSet,
    io::{self},
};

use crate::utils::read_lines;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day11/input_{run_as}.txt");

    let lines = read_lines(&input_file);

    let mut space_map = lines
        .clone()
        .iter()
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut new_y_indexes = vec![];
    for y in 0..space_map.len() {
        if space_map[y].iter().all(|&x| x == '.') {
            new_y_indexes.push(y);
        }
    }

    let mut offset = 0;
    for index in &new_y_indexes {
        space_map.insert(index + offset, vec!['.'; space_map[0].len()]);
        offset += 1;
    }

    let mut new_x_indexes = vec![];
    for x in 0..space_map[0].len() {
        if space_map.iter().all(|y| y[x] == '.') {
            new_x_indexes.push(x);
        }
    }

    offset = 0;
    for index in &new_x_indexes {
        for y in 0..space_map.len() {
            space_map[y].insert(index + offset, '.');
        }
        offset += 1;
    }

    // print_space_map(space_map.clone());

    let mut count = part1(space_map) as i64;
    println!("Part 1: {}", count);

    space_map = lines
        .clone()
        .iter()
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    count = part2(space_map, new_x_indexes, new_y_indexes);
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(space_map: Vec<Vec<char>>) -> i32 {
    let mut result = 0;
    let mut galaxies = vec![];

    for y in 0..space_map.len() {
        for x in 0..space_map[0].len() {
            if space_map[y][x] == '#' {
                galaxies.push((x as i32, y as i32));
            }
        }
    }

    // using a Hashset is really important here, as the list gets longer and longer
    let mut visited_paths = HashSet::new();
    for i in 0..galaxies.len() {
        for j in 0..galaxies.len() {
            if j == i {
                continue;
            }

            if visited_paths.contains(&(i, j)) || visited_paths.contains(&(j, i)) {
                continue;
            }

            visited_paths.insert((i, j));

            result += (galaxies[i].0 - galaxies[j].0).abs() + (galaxies[i].1 - galaxies[j].1).abs();
        }
    }

    result
}

fn part2(space_map: Vec<Vec<char>>, x_indexes: Vec<usize>, y_indexes: Vec<usize>) -> i64 {
    let mut result: i64 = 0;
    let mut galaxies = vec![];

    for y in 0..space_map.len() {
        for x in 0..space_map[0].len() {
            if space_map[y][x] == '#' {
                galaxies.push((x as i32, y as i32));
            }
        }
    }

    let mut expansions: i64 = 0;
    // using a Hashset is really important here, as the list gets longer and longer
    let mut visited_paths = HashSet::new();
    for i in 0..galaxies.len() {
        for j in 0..galaxies.len() {
            if j == i {
                continue;
            }

            if visited_paths.contains(&(i, j)) || visited_paths.contains(&(j, i)) {
                continue;
            }

            visited_paths.insert((i, j));

            let min_x = cmp::min(galaxies[i].0, galaxies[j].0);
            let max_x = cmp::max(galaxies[i].0, galaxies[j].0);

            let x_expansions = x_indexes
                .iter()
                .filter(|&&x| (x as i32) > min_x && (x as i32) < max_x)
                .count();

            expansions += x_expansions as i64;

            let min_y = cmp::min(galaxies[i].1, galaxies[j].1);
            let max_y = cmp::max(galaxies[i].1, galaxies[j].1);

            let y_expansions = y_indexes
                .iter()
                .filter(|&&y| (y as i32) > min_y && (y as i32) < max_y)
                .count();

            expansions += y_expansions as i64;

            result += ((galaxies[i].0 - galaxies[j].0).abs()
                + (galaxies[i].1 - galaxies[j].1).abs()) as i64;
        }
    }

    result + expansions * (1_000_000 - 1) // - 1, because the result already counted the row/column that expands
}

fn print_space_map(space_map: Vec<Vec<char>>) {
    for y in 0..space_map.len() {
        for x in 0..space_map[0].len() {
            print!("{}", space_map[y][x]);
        }
        println!();
    }
}
