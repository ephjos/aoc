/*
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
*/

fn rewrite(input: &str) {
    let lines = input
        .lines()
        .map(|l| {
            l.parse::<usize>()
                .expect(&format!("Could not convert line {} to usize", l))
        })
        .collect::<Vec<usize>>();
    let mut prev = usize::MAX;
    let mut count_01 = 0;
    let mut count_02 = 0;

    for i in 0..3 {
        let curr = lines[i];
        if curr > prev {
            count_01 += 1;
        }
        prev = curr;
    }

    for i in 3..lines.len() {
        let curr = lines[i];
        if curr > prev {
            count_01 += 1;
        }
        prev = curr;

        // Since the windows are overlapping, only compare the new (rightmost)
        // and old (outside to the left of the window) values
        if curr > lines[i-3] {
            count_02 += 1;
        }
    }
    println!("01.1: {:?}", count_01);
    println!("01.2: {:?}", count_02);
}

pub fn run() {
    let input = include_str!("../input/day01");
    // println!("01.1: {:?}", part1(include_str!("../input/day01")));
    // println!("01.2: {:?}", part2(include_str!("../input/day01")));
    rewrite(input);
}
