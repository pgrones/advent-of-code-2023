use regex::Regex;
use std::{cmp::Ordering, collections::HashMap, io};

use crate::utils::{read_lines_into_structs, Instantiable};

#[derive(Debug, Clone, PartialEq)]
enum WinType {
    FiveOfAKind = 6,
    FourOfAKind = 5,
    FullHouse = 4,
    ThreeOfAKind = 3,
    TwoPair = 2,
    OnePair = 1,
    HighCard = 0,
}

#[derive(Debug, Clone)]
struct Hand {
    cards: String,
    bid: u32,
}

impl Hand {
    fn win_type(&self) -> WinType {
        let mut tokens: HashMap<char, i8> = HashMap::new();

        for card in self.cards.chars().collect::<Vec<char>>() {
            *tokens.entry(card).or_insert(0) += 1;
        }

        if tokens.len() == 1 {
            return WinType::FiveOfAKind;
        }

        let mut values: Vec<&i8> = tokens.values().collect();
        values.sort_unstable();

        if values.len() == 2 {
            let score = values[1] - values[0];

            if score == 3 {
                return WinType::FourOfAKind;
            } else {
                return WinType::FullHouse;
            }
        }

        if values.len() == 3 {
            let score = values[2] - values[1] - values[0];

            if score == 1 {
                return WinType::ThreeOfAKind;
            } else {
                return WinType::TwoPair;
            }
        }

        if values.len() == 4 {
            return WinType::OnePair;
        }

        WinType::HighCard
    }
}

impl Instantiable for Hand {
    fn new(values: Vec<&str>) -> Self {
        Hand {
            cards: values[0].to_string(),
            bid: values[1].parse::<u32>().unwrap(),
        }
    }
}

const STRENGTHS: [char; 13] = [
    'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2',
];

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day7/input_{run_as}.txt");

    let hands =
        read_lines_into_structs::<Hand, _>(&input_file, Regex::new(r"\s").unwrap(), |_, _| true);

    let mut result = part1(hands.clone());
    println!("Part 1: {}", result);

    // result = part2(seeds, maps);
    // println!("Part 2: {}", result);

    Ok(())
}

fn part1(mut hands: Vec<Hand>) -> u32 {
    hands.sort_by(|a, b| {
        let ord = (a.win_type() as u8).cmp(&(b.win_type() as u8));

        if ord != Ordering::Equal {
            return ord;
        }

        let a_chars: Vec<char> = a.cards.chars().collect();
        let b_chars: Vec<char> = b.cards.chars().collect();

        for i in 0..a_chars.len() {
            if a_chars[i] != b_chars[i] {
                let a_strength = STRENGTHS.iter().position(|&x| x == a_chars[i]).unwrap();
                let b_strength = STRENGTHS.iter().position(|&x| x == b_chars[i]).unwrap();

                // a is better when its index is smaller
                if a_strength < b_strength {
                    return Ordering::Greater;
                } else {
                    return Ordering::Less;
                }
            }
        }

        Ordering::Equal
    });

    hands
        .iter()
        .enumerate()
        .fold(0, |acc, (i, x)| acc + x.bid * ((i as u32) + 1))
}

// fn part2(seeds: Vec<i64>, maps: Vec<Vec<Instruction>>) -> u32 {
//     let mut result = 0;
//     result
// }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_win_types() {
        assert_eq!(
            WinType::FiveOfAKind,
            Hand::new(vec!("AAAAA", "1")).win_type()
        );

        assert_eq!(
            WinType::FourOfAKind,
            Hand::new(vec!("AA8AA", "1")).win_type()
        );

        assert_eq!(WinType::FullHouse, Hand::new(vec!("23332", "1")).win_type());

        assert_eq!(
            WinType::ThreeOfAKind,
            Hand::new(vec!("TTT98", "1")).win_type()
        );

        assert_eq!(WinType::TwoPair, Hand::new(vec!("23432", "1")).win_type());

        assert_eq!(WinType::OnePair, Hand::new(vec!("A23A4", "1")).win_type());

        assert_eq!(WinType::HighCard, Hand::new(vec!("23456", "1")).win_type());
    }
}
