use std::{
    cmp,
    io::{self},
    u32::MAX,
    vec,
};

use crate::utils::read_lines;

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day17/input_{run_as}.txt");

    let city_map: Vec<Vec<u32>> = read_lines(&input_file)
        .iter()
        .map(|x| x.chars().map(|y| y.to_digit(10).unwrap()).collect())
        .collect();

    let mut count = part1(city_map);
    println!("Part 1: {}", count);

    count = part2();
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(city_map: Vec<Vec<u32>>) -> u32 {
    let mut min_dists = vec![];

    for _ in 0..city_map.len() {
        let mut row = vec![];
        for _ in 0..city_map[0].len() {
            row.push((MAX, (0, 0)));
        }
        min_dists.push(row);
    }

    let mut radius = 1;
    let size_x = city_map[0].len();
    let size_y = city_map.len();
    let max_radius = cmp::max(size_x, size_y);

    min_dists[size_y - 1][size_x - 1] = (city_map[size_y - 1][size_x - 1], (1000, 1000));

    while radius < max_radius {
        let mut coords = vec![];

        for y in (size_y - radius - 1..size_y).rev() {
            coords.push((size_x - radius - 1, y));
        }

        for x in size_x - radius..size_x {
            coords.push((x, size_y - radius - 1));
        }

        for (index, coord) in coords.iter().enumerate() {
            min_dists[coord.1][coord.0] =
                computemagic(&min_dists, &city_map, &coords, index, radius);
        }

        radius += 1;
    }

    print_map(&min_dists);

    min_dists[0][0].0 - city_map[0][0]
}

fn computemagic(
    min_dists: &Vec<Vec<(u32, (usize, usize))>>,
    city_map: &Vec<Vec<u32>>,
    coords: &Vec<(usize, usize)>,
    index: usize,
    radius: usize,
) -> (u32, (usize, usize)) {
    let mut global_min = MAX;
    let mut successor = (1000000000, 1000000000);
    let corner_index = radius;

    for i in 0..2 * radius + 1 {
        if index + i < radius || index + i > coords.len() + radius - 1 {
            continue;
        }

        let curr_index = index + i - radius;

        let neighbor_dist;

        if curr_index < corner_index {
            neighbor_dist = min_dists[coords[curr_index].1][coords[curr_index].0 + 1];
        } else if curr_index > corner_index {
            neighbor_dist = min_dists[coords[curr_index].1 + 1][coords[curr_index].0];
        } else {
            continue;
        }

        let sum: &u32 = &coords[cmp::min(curr_index, index)..cmp::max(curr_index, index) + 1]
            .iter()
            .map(|(x, y)| city_map[*y][*x])
            .sum();

        if sum + neighbor_dist.0 < global_min {
            global_min = sum + neighbor_dist.0;
            if sum > &city_map[coords[index].1][coords[index].0] {
                if i < radius {
                    successor = coords[index - 1];
                } else {
                    successor = coords[index + 1];
                }
            } else {
                if curr_index < corner_index {
                    successor = (coords[curr_index].0 + 1, coords[curr_index].1);
                } else if curr_index > corner_index {
                    successor = (coords[curr_index].0, coords[curr_index].1 + 1);
                }
            }
        }
    }

    (global_min, successor)
}

fn part2() -> u32 {
    0
}

fn print_map(map: &Vec<Vec<(u32, (usize, usize))>>) {
    let mut path = vec![];
    let mut coord = map[0][0].1;
    while coord != (1000, 1000) {
        path.push(coord);
        coord = map[coord.1][coord.0].1;
    }

    for y in 0..map.len() {
        for x in 0..map[0].len() {
            if path.contains(&(x, y)) {
                print!(
                    "\x1b[93m{} \x1b[0m",
                    format!("{:width$}", map[y][x].0, width = 3)
                );
            } else {
                print!("{} ", format!("{:width$}", map[y][x].0, width = 3));
            }
        }
        println!();
    }
}
