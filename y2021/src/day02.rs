/*
fn part1(input: &str) -> usize {
    let mut x = 0;
    let mut depth = 0;

    for line in input.lines() {
        let toks = line.split(" ").collect::<Vec<_>>();
        let cmd = toks[0];
        let val = toks[1].parse::<usize>().expect("Could not parse value");

        match cmd {
            "up" => depth -= val,
            "down" => depth += val,
            "forward" => x += val,
            "forward" => x += val,
            _ => (),
        }
    }
    return x * depth;
}

fn part2(input: &str) -> usize {
    let mut x = 0;
    let mut depth = 0;
    let mut aim = 0;

    for line in input.lines() {
        let toks = line.split(" ").collect::<Vec<_>>();
        let cmd = toks[0];
        let val = toks[1].parse::<usize>().expect("Could not parse value");

        match cmd {
            "up" => aim -= val,
            "down" => aim += val,
            "forward" => {
                x += val;
                depth += aim * val;
            },
            _ => (),
        }
    }
    return x * depth;
}
*/

fn rewrite(input: &str) {
    let mut x = 0;
    let mut depth1 = 0;
    let mut depth2 = 0;
    let mut aim = 0;

    for line in input.lines() {
        let toks = line.splitn(2, " ").collect::<Vec<_>>();
        let cmd = toks[0].chars().nth(0).expect("Could not get command");
        let val = toks[1].parse::<usize>().expect("Could not parse value");

        match cmd {
            'u' => {
                aim -= val;
                depth1 -= val;
            }
            'd' => {
                aim += val;
                depth1 += val;
            }
            'f' => {
                x += val;
                depth2 += aim * val;
            }
            _ => unreachable!(),
        }
    }
    println!("02.1: {:?}", x * depth1);
    println!("02.1: {:?}", x * depth2);
}

pub fn run() {
    let input = include_str!("../input/day02");
    //println!("02.1: {:?}", part1(input));
    //println!("02.2: {:?}", part2(input));
    rewrite(input);
}
