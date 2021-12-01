
fn part1(input: &str) -> isize {
    let mut prev = usize::MAX;
    let mut count = 0;
    for line in input.lines() {
        let curr = line.parse::<usize>().expect(&format!("Could not convert {}", line));
        if curr > prev {
            count += 1;
        }
        prev = curr;
    }
    return count;
}

fn part2(input: &str) -> isize {
    let lines = input.lines().map(|l| l.parse::<usize>().expect(&format!("Could not convert line {} to usize", l))).collect::<Vec<usize>>();
    let mut prev = usize::MAX;
    let mut count = 0;

    for i in 2..lines.len() {
        let curr = lines[i-2] + lines[i-1] + lines[i];
        if curr > prev {
            count += 1;
        }
        prev = curr;
    }
    return count;
}

pub fn run() {
    println!("01.1: {:?}", part1(include_str!("../input/day01")));
    println!("01.2: {:?}", part2(include_str!("../input/day01")));
}

