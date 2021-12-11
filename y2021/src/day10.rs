use std::collections::HashMap;
use std::collections::VecDeque;

fn part1(input: &str) -> i64 {
    let mut stack: VecDeque<char> = VecDeque::new();
    let mut points = 0;

    let point_map: HashMap<char,i64> = HashMap::from([
        (')', 3),
        (']', 57),
        ('}', 1197),
        ('>', 25137),
    ]);

    for line in input.trim().lines() {
        for c in line.trim().chars() {
            match c {
                '(' => {
                    stack.push_front(')');
                },
                '[' => {
                    stack.push_front(']');
                },
                '{' => {
                    stack.push_front('}');
                },
                '<' => {
                    stack.push_front('>');
                },
                _ => {
                    if let Some(top) = stack.pop_front() {
                        if c != top {
                            points += point_map[&c];
                        }
                    }
                }
            }
        }
    }

    return points;
}

fn part2(input: &str) -> i64 {
    let point_map: HashMap<char,i64> = HashMap::from([
        (')', 3),
        (']', 57),
        ('}', 1197),
        ('>', 25137),
    ]);
    let autocomplete_point_map: HashMap<char,i64> = HashMap::from([
        (')', 1),
        (']', 2),
        ('}', 3),
        ('>', 4),
    ]);
    let mut autocomplete_points: Vec<i64> = Vec::new();

    for line in input.trim().lines() {
        let mut stack: VecDeque<char> = VecDeque::new();
        let mut points = 0;
        for c in line.trim().chars() {
            match c {
                '(' => {
                    stack.push_front(')');
                },
                '[' => {
                    stack.push_front(']');
                },
                '{' => {
                    stack.push_front('}');
                },
                '<' => {
                    stack.push_front('>');
                },
                _ => {
                    if let Some(top) = stack.pop_front() {
                        if c != top {
                            points += point_map[&c];
                        }
                    }
                }
            }
        }

        if points != 0 {
            // Skip corrupted lines
            continue;
        }

        let mut score = 0;
        for c in stack.iter() {
            score *= 5;
            score += autocomplete_point_map[&c];
        }
        // TODO: use a binary tree, skip sort?
        autocomplete_points.push(score);
    }

    autocomplete_points.sort();
    return autocomplete_points[autocomplete_points.len() / 2];
}

pub fn run() {
    let input = include_str!("../input/day10");
    println!("10.1: {:?}", part1(input));
    println!("10.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]"), 26397);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]"), 288957);
    }
}
