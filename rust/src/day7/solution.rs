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

impl Instantiable for Hand {
    fn new(values: Vec<&str>) -> Self {
        Hand {
            cards: values[0].to_string(),
            bid: values[1].parse::<u32>().unwrap(),
        }
    }
}

impl Hand {
    fn win_type(&self, treat_j_as_joker: bool) -> WinType {
        // count each occurrence of a card in the hand
        let mut cards: HashMap<char, i8> = HashMap::new();

        for card in self.cards.chars().collect::<Vec<char>>() {
            *cards.entry(card).or_insert(0) += 1;
        }

        // for part 2, j is considered a joker and takes the form of the best possible card
        // the best possible card is always going to be the most common one
        if treat_j_as_joker {
            if cards.len() > 1 && cards.contains_key(&'J') {
                let j_amount = cards.remove(&'J').unwrap();

                let best_card = cards
                    .iter()
                    .max_by(|a, b| a.1.cmp(&b.1))
                    .map(|(k, _v)| k)
                    .unwrap();

                cards.entry(*best_card).and_modify(|x| *x += j_amount);
            }
        }

        // all cards are the same
        if cards.len() == 1 {
            return WinType::FiveOfAKind;
        }

        let mut values: Vec<&i8> = cards.values().collect();
        values.sort_unstable();

        // two unique cards
        if values.len() == 2 {
            // generate a pseudo-hash to identify the type
            let pseudo_hash = values[1] - values[0];

            if pseudo_hash == 3 {
                return WinType::FourOfAKind;
            } else {
                return WinType::FullHouse;
            }
        }

        //  three unique cards
        if values.len() == 3 {
            // generate a pseudo-hash to identify the type
            let pseudo_hash = values[2] - values[1] - values[0];

            if pseudo_hash == 1 {
                return WinType::ThreeOfAKind;
            } else {
                return WinType::TwoPair;
            }
        }

        // four unique cards
        if values.len() == 4 {
            return WinType::OnePair;
        }

        WinType::HighCard
    }
}

pub fn solve(run_as: char) -> Result<(), io::Error> {
    let input_file = format!("src/day7/input_{run_as}.txt");

    let hands =
        read_lines_into_structs::<Hand, _>(&input_file, Regex::new(r"\s").unwrap(), |_, _| true);

    let mut strengths: [char; 13] = [
        'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2',
    ];

    let mut result = part1_and_2(hands.clone(), strengths, false);
    println!("Part 1: {}", result);

    strengths = [
        'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J',
    ];

    result = part1_and_2(hands, strengths, true);
    println!("Part 2: {}", result);

    Ok(())
}

fn part1_and_2(mut hands: Vec<Hand>, strengths: [char; 13], treat_j_as_joker: bool) -> u32 {
    hands.sort_by(|a, b| {
        // different win types always have a clear winner
        let ord = (a.win_type(treat_j_as_joker) as u8).cmp(&(b.win_type(treat_j_as_joker) as u8));

        if ord != Ordering::Equal {
            return ord;
        }

        // if the win types are the same, compare the cards individually
        let a_chars: Vec<char> = a.cards.chars().collect();
        let b_chars: Vec<char> = b.cards.chars().collect();

        for i in 0..a_chars.len() {
            if a_chars[i] != b_chars[i] {
                let a_strength = strengths.iter().position(|&x| x == a_chars[i]).unwrap();
                let b_strength = strengths.iter().position(|&x| x == b_chars[i]).unwrap();

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

    // reduce the bids down to a single number
    hands
        .iter()
        .enumerate()
        .fold(0, |acc, (i, x)| acc + x.bid * ((i as u32) + 1))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_win_types() {
        assert_eq!(
            WinType::FiveOfAKind,
            Hand::new(vec!("AAAAA", "1")).win_type(false)
        );

        assert_eq!(
            WinType::FourOfAKind,
            Hand::new(vec!("AA8AA", "1")).win_type(false)
        );

        assert_eq!(
            WinType::FullHouse,
            Hand::new(vec!("23332", "1")).win_type(false)
        );

        assert_eq!(
            WinType::ThreeOfAKind,
            Hand::new(vec!("TTT98", "1")).win_type(false)
        );

        assert_eq!(
            WinType::TwoPair,
            Hand::new(vec!("23432", "1")).win_type(false)
        );

        assert_eq!(
            WinType::OnePair,
            Hand::new(vec!("A23A4", "1")).win_type(false)
        );

        assert_eq!(
            WinType::HighCard,
            Hand::new(vec!("23456", "1")).win_type(false)
        );
    }
}
