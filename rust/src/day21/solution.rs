use std::{
    collections::HashSet,
    io::{self},
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
    count = part2(&garden);
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(garden: &Vec<Vec<char>>) -> u64 {
    let mut start = (0, 0);

    'outer: for y in 0..garden.len() {
        for x in 0..garden[0].len() {
            if garden[y][x] == 'S' {
                start = (x, y);
                break 'outer;
            }
        }
    }

    let mut queue: Vec<((usize, usize), bool)> = vec![(start, false)];

    for _ in 0..64 {
        let mut next_positions = HashSet::new();

        while queue.len() > 0 {
            next_positions.extend(execute_step(garden, queue.pop().unwrap(), true));
        }

        queue = Vec::from_iter(next_positions);
    }

    return queue
        .iter()
        .filter(|((x, y), a)| !a && y.abs_diff(65) + x.abs_diff(65) <= 65)
        .count() as u64;
}

fn execute_step(
    garden: &Vec<Vec<char>>,
    curr_pos: ((usize, usize), bool),
    wrap: bool,
) -> Vec<((usize, usize), bool)> {
    let mut result = vec![];

    if curr_pos.0 .1 > 0 && garden[curr_pos.0 .1 - 1][curr_pos.0 .0] != '#' {
        result.push(((curr_pos.0 .0, curr_pos.0 .1 - 1), curr_pos.1));
    } else if curr_pos.0 .1 == 0 && wrap {
        result.push(((curr_pos.0 .0, garden.len() - 1), !curr_pos.1));
    }

    if curr_pos.0 .1 < garden.len() - 1 && garden[curr_pos.0 .1 + 1][curr_pos.0 .0] != '#' {
        result.push(((curr_pos.0 .0, curr_pos.0 .1 + 1), curr_pos.1));
    } else if curr_pos.0 .1 == garden.len() - 1 && wrap {
        result.push(((curr_pos.0 .0, 0), !curr_pos.1));
    }

    if curr_pos.0 .0 > 0 && garden[curr_pos.0 .1][curr_pos.0 .0 - 1] != '#' {
        result.push(((curr_pos.0 .0 - 1, curr_pos.0 .1), curr_pos.1));
    } else if curr_pos.0 .0 == 0 && wrap {
        result.push(((garden[0].len() - 1, curr_pos.0 .1), !curr_pos.1));
    }

    if curr_pos.0 .0 < garden[0].len() - 1 && garden[curr_pos.0 .1][curr_pos.0 .0 + 1] != '#' {
        result.push(((curr_pos.0 .0 + 1, curr_pos.0 .1), curr_pos.1));
    } else if curr_pos.0 .0 == garden[0].len() - 1 && wrap {
        result.push(((0, curr_pos.0 .1), !curr_pos.1));
    }

    result
}

// fn render_garden(garden: &Vec<Vec<char>>, plots: &Vec<(usize, usize)>, multiplicity: usize) {
//     for y in 0..multiplicity * garden.len() {
//         for x in 0..multiplicity * garden[0].len() {
//             if plots.contains(&(x % garden.len(), y % garden[0].len())) {
//                 print!("O");
//             } else {
//                 print!("{}", garden[y % garden[0].len()][x % garden.len()]);
//             }
//         }
//         println!();
//     }
//     println!();
// }

const STEPS: u64 = 26501365;

fn part2(garden: &Vec<Vec<char>>) -> u64 {
    let center = garden.len() / 2;

    let remaining_steps = (STEPS - center as u64) % garden.len() as u64 + garden.len() as u64;
    let edge = (STEPS - center as u64) / garden.len() as u64 - 1;

    let edge_length = 1 + edge * 2;

    // let number_of_gardens = edge_length * edge_length / 2;
    // let number_of_inside_diamonds = number_of_gardens + 1;
    // let number_of_outside_diamonds = number_of_gardens;

    // 3632
    // 3701

    let mut diamond_indexes = vec![];
    let mut number_of_plots = 0;
    let mut number_of_plots_inside_diamond = 0;
    let mut number_of_plots_odd = 0;
    let mut number_of_plots_inside_diamond_odd = 0;

    for y in 0..garden.len() {
        for x in 0..garden[0].len() {
            if y.abs_diff(center) + x.abs_diff(center) == center {
                diamond_indexes.push((x, y));
            }

            if (x + y) % 2 == 1 && (garden[y][x] == '.' || garden[y][x] == 'S') {
                number_of_plots += 1;

                if y.abs_diff(center) + x.abs_diff(center) <= center {
                    number_of_plots_inside_diamond += 1;
                }
            }

            if (x + y) % 2 == 0 && (garden[y][x] == '.' || garden[y][x] == 'S') {
                number_of_plots_odd += 1;

                if y.abs_diff(center) + x.abs_diff(center) <= center {
                    number_of_plots_inside_diamond_odd += 1;
                }
            }
        }
    }

    println!("{}", number_of_plots_inside_diamond);

    let number_of_plots_outside_diamond = number_of_plots - number_of_plots_inside_diamond;
    let number_of_plots_outside_diamond_odd =
        number_of_plots_odd - number_of_plots_inside_diamond_odd;

    let mut odd_multiplier = 1;

    for i in (5..edge_length + 1).step_by(4) {
        odd_multiplier += 4 * (i - 1);
    }

    let mut even_multiplier = 0;

    for i in (3..edge_length + 1).step_by(4) {
        even_multiplier += 4 * (i - 1);
    }

    let number_of_plots_inside_fat_ass_diamond = even_multiplier / 2
        * number_of_plots_inside_diamond
        + (odd_multiplier / 2 + 1) * number_of_plots_inside_diamond_odd
        + even_multiplier / 2 * number_of_plots_outside_diamond
        + odd_multiplier / 2 * number_of_plots_outside_diamond_odd;

    let mut queue: Vec<((usize, usize), bool)> =
        diamond_indexes.iter().map(|x| (*x, false)).collect();

    for _ in 0..remaining_steps {
        let mut next_positions = HashSet::new();

        while queue.len() > 0 {
            next_positions.extend(execute_step(garden, queue.pop().unwrap(), true));
        }

        queue = Vec::from_iter(next_positions)
    }

    let reachable = queue
        .iter()
        .filter(|&((x, y), wrapped)| *wrapped || y.abs_diff(center) + x.abs_diff(center) > center)
        .map(|x| x.0)
        .count() as u64;

    let result = number_of_plots_inside_fat_ass_diamond + reachable * edge_length;

    println!(
        "number_of_plots_inside_fat_ass_diamond: {}",
        number_of_plots_inside_fat_ass_diamond
    );
    println!("reachable: {}", reachable);
    println!("edge_length: {}", edge_length);
    println!("reachable * edge_length: {}", reachable * edge_length);

    result
}

// 600336060511101 (correct)
// 600661970655197
// 596405782573597
// 602094355805197
// 599231058260303
//      4465559163
//     81850580000
