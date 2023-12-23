use std::{collections::HashMap, io};

use crate::utils::{char_to_ascii, read_lines};

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day15/input_{run_as}.txt");

    let init_sequence = read_lines(&input_file)[0]
        .split(",")
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut count = part1(&init_sequence);
    println!("Part 1: {}", count);

    count = part2(&init_sequence);
    println!("Part 2: {}", count);

    Ok(())
}

fn part1(init_sequence: &Vec<Vec<char>>) -> u32 {
    init_sequence
        .iter()
        .fold(0, |acc, x| acc + x.iter().fold(0, |acc, &c| hash(acc, c)))
}

fn part2(init_sequence: &Vec<Vec<char>>) -> u32 {
    let mut boxes: Vec<HashMap<String, (u32, usize)>> = vec![];
    for _ in 0..256 {
        boxes.push(HashMap::new());
    }

    for (i, step) in init_sequence.iter().enumerate() {
        let label = &step[0..step.iter().position(|x| ['=', '-'].contains(x)).unwrap()];
        let operation = if step.contains(&'=') {
            Operation::Add
        } else {
            Operation::Remove
        };
        let focal_length = step.last().unwrap();

        let hash = label.iter().fold(0, |acc, &c| hash(acc, c));

        let map = &mut boxes[hash as usize];

        let label_as_string: String = label.iter().collect();

        if operation == Operation::Add {
            if map.contains_key(&label_as_string) {
                map.get_mut(&label_as_string).unwrap().0 = focal_length.to_digit(10).unwrap();
            } else {
                map.insert(label_as_string, (focal_length.to_digit(10).unwrap(), i));
            }
        } else {
            map.remove(&label_as_string);
        }
    }

    let mut result = 0;

    for (i, box_) in boxes.iter().enumerate() {
        if box_.is_empty() {
            continue;
        }

        let mut lenses = box_.values().collect::<Vec<&(u32, usize)>>();
        lenses.sort_by(|&a, &b| a.1.cmp(&b.1));

        result += lenses.iter().enumerate().fold(0, |acc, (j, &x)| {
            acc + (i as u32 + 1) * (j as u32 + 1) * x.0
        })
    }

    result
}

fn hash(mut curr_value: u32, c: char) -> u32 {
    // Determine the ASCII code for the current character of the string.
    let ascii_code = char_to_ascii(c) as u32;
    // Increase the current value by the ASCII code you just determined.
    curr_value += ascii_code;
    // Set the current value to itself multiplied by 17.
    curr_value *= 17;
    // Set the current value to the remainder of dividing itself by 256.
    curr_value %= 256;
    // The current value is the output of the HASH algorithm
    curr_value
}

#[derive(PartialEq)]
enum Operation {
    Add,
    Remove,
}
