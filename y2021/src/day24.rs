
//
// This solution is entirely based https://github.com/dphilipson/advent-of-code-2021/blob/master/src/days/day24.rs
// This write-up was phenomenal and outlines a simple method for solving this problem.
// I initially tried it by hand and was running into issues, so I figured that
// I could try to implement it and see if I found the problem. This helped
// me figure out what I was doing wrong, and it was fun to implement this
// approach.
//

struct Entry {
    i: usize,
    v: isize,
}

#[derive(Debug)]
struct Def {
    l: usize, // Index of left digit
    r: usize, // Index of right digit
    v: isize, // Value added to right digit
}

fn part1(input: &str) -> isize {
    let mut stack: Vec<Entry>  = Vec::new();
    let blocks = input.trim().split("inp w\n").collect::<Vec<&str>>();

    let check_offset_pairs = blocks[1..].iter().map(|b| {
        let lines = b.lines().collect::<Vec<&str>>();
        let check_parts = lines[4].split(" ").collect::<Vec<&str>>();
        let offset_parts = lines[14].split(" ").collect::<Vec<&str>>();
        return (check_parts[2].parse::<isize>().unwrap(), offset_parts[2].parse::<isize>().unwrap());
    }).collect::<Vec<(isize, isize)>>();

    let mut defs: Vec<Def> = Vec::new();
    for (i, (check, offset)) in check_offset_pairs.iter().enumerate() {
        if *check > 1 {
            stack.push(Entry {
                i,
                v: *offset,
            });
        } else {
            let popped = stack.pop().unwrap();
            let v = check + popped.v;
            if v > 0 {
                defs.push(Def {
                    l: popped.i,
                    r: i,
                    v: -v,
                });
            } else {
                defs.push(Def {
                    l: i,
                    r: popped.i,
                    v,
                });
            }
        }
    }

    let mut res: [isize; 14] = [0; 14];

    for def in defs {
        let l = def.l;
        let r = def.r;
        let v = def.v;
        // Print out the rules in the format used in the source material
        // println!("input[{}] = input[{}] + {}", l, r, v);

        res[l] = 9+v;
        res[r] = 9;
    }

    return res.iter().fold(0, |acc, ele| {
        let temp = acc * 10;
        return temp + ele;
    });
}

fn part2(input: &str) -> isize {
    let mut stack: Vec<Entry>  = Vec::new();
    let blocks = input.trim().split("inp w\n").collect::<Vec<&str>>();

    let check_offset_pairs = blocks[1..].iter().map(|b| {
        let lines = b.lines().collect::<Vec<&str>>();
        let check_parts = lines[4].split(" ").collect::<Vec<&str>>();
        let offset_parts = lines[14].split(" ").collect::<Vec<&str>>();
        return (check_parts[2].parse::<isize>().unwrap(), offset_parts[2].parse::<isize>().unwrap());
    }).collect::<Vec<(isize, isize)>>();

    let mut defs: Vec<Def> = Vec::new();
    for (i, (check, offset)) in check_offset_pairs.iter().enumerate() {
        if *check > 1 {
            stack.push(Entry {
                i,
                v: *offset,
            });
        } else {
            let popped = stack.pop().unwrap();
            let v = check + popped.v;
            if v > 0 {
                defs.push(Def {
                    l: popped.i,
                    r: i,
                    v: -v,
                });
            } else {
                defs.push(Def {
                    l: i,
                    r: popped.i,
                    v,
                });
            }
        }
    }

    let mut res: [isize; 14] = [0; 14];

    for def in defs {
        let l = def.l;
        let r = def.r;
        let v = def.v;
        // Print out the rules in the format used in the source material
        // println!("input[{}] = input[{}] + {}", l, r, v);

        res[l] = 1;
        res[r] = 1-v;
    }

    return res.iter().fold(0, |acc, ele| {
        let temp = acc * 10;
        return temp + ele;
    });
}

pub fn run() {
    let input = include_str!("../input/day24");
    println!("24.1: {:?}", part1(input));
    println!("24.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(""), 1);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(""), 1);
    }
}
