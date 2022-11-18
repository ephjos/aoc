use std::fs;

#[derive(Debug)]
enum Op {
    Toggle,
    On,
    Off,
}

#[derive(Debug)]
struct Point2D {
    x: usize,
    y: usize,
}

#[derive(Debug)]
struct Instr {
    op: Op,
    low: Point2D,
    high: Point2D,
}

fn parse_instr(s: &str) -> Instr {
    let s = s.trim();
    let mut toks = s.split(" ").collect::<Vec<_>>();
    let mut op: Op = Op::Toggle;

    match toks[0] {
        "toggle" => op = Op::Toggle,
        "turn" => {
            if toks[1] == "on" {
                op = Op::On;
            } else {
                op = Op::Off;
            }
            toks.remove(0);
        },
        _ => (),
    }

    let ns = toks[1].split(",")
        .map(|n| n.parse::<usize>().unwrap())
        .collect::<Vec<_>>();
    let low = Point2D { x: ns[0], y: ns[1] };
    let ns = toks[3].split(",")
        .map(|n| n.parse::<usize>().unwrap())
        .collect::<Vec<_>>();
    let high = Point2D { x: ns[0], y: ns[1] };

    return Instr {
        op,
        low,
        high,
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
enum Light {
    On,
    Off,
}

const GRID_N: usize = 1000;

fn a (instrs: &Vec<Instr>) -> usize {
    let mut lights = [[Light::Off; GRID_N]; GRID_N];

    for instr in instrs {
        for i in instr.low.x..=instr.high.x {
            for j in instr.low.y..=instr.high.y {
                match instr.op {
                    Op::On => lights[i][j] = Light::On,
                    Op::Off => lights[i][j] = Light::Off,
                    Op::Toggle => lights[i][j] = if lights[i][j] == Light::On {
                        Light::Off
                    } else {
                        Light::On
                    },
                }
            }
        }
    }

    return lights.iter().flatten().map(|l| (*l == Light::On) as usize).sum();
}

fn b (instrs: &Vec<Instr>) -> usize {
    let mut lights = [[0; GRID_N]; GRID_N];

    for instr in instrs {
        for i in instr.low.x..=instr.high.x {
            for j in instr.low.y..=instr.high.y {
                match instr.op {
                    Op::On => lights[i][j] += 1,
                    Op::Off => {
                        lights[i][j] = lights[i][j]-(lights[i][j] != 0) as usize;
                    },
                    Op::Toggle => lights[i][j] += 2,
                }
            }
        }
    }

    return lights.iter().flatten().sum();
}

y2015::main! {
    let content = fs::read_to_string("./input/day06.txt").unwrap();
    let instrs = content.lines()
        .map(|line| parse_instr(line))
        .collect::<Vec<_>>();

    (a(&instrs), b(&instrs))
}
