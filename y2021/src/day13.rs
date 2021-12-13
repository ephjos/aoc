use std::collections::HashSet;

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
struct Point2d {
    x: u32,
    y: u32,
}

#[derive(Debug, Copy, Clone, PartialEq, Eq)]
struct Fold {
    axis: char,
    line: u32,
}

fn part1(input: &str) -> usize {
    let sections: Vec<&str> = input.trim().splitn(2, "\n\n").collect();
    let dots: HashSet<Point2d> = sections[0].lines().map(|l| {
        let vs: Vec<u32> = l.splitn(2, ",").map(|x| x.parse::<u32>().unwrap()).collect();
        return Point2d { x: vs[0], y: vs[1] };
    }).collect();
    let folds: Vec<Fold> = sections[1].lines().map(|l| {
        let end: &str = l.trim().split(" ").collect::<Vec<&str>>()[2];
        let vs: Vec<&str> = end.splitn(2, "=").collect();
        return Fold {
            axis: vs[0].chars().next().unwrap(),
            line: vs[1].parse::<u32>().unwrap(),
        };
    }).collect();

    let axis = folds[0].axis;
    let line = folds[0].line;
    let dots_after_fold_0: HashSet<Point2d> = dots.iter().map(|d| {
        if axis == 'y' {
            if d.y > line {
                return Point2d {
                    x: d.x,
                    y: line - (d.y - line),
                };
            } else {
                return *d;
            }
        }

        if d.x > line {
            return Point2d {
                x: line - (d.x - line),
                y: d.y,
            };
        } else {
            return *d;
        }
    }).collect();

    return dots_after_fold_0.len();
}

fn part2(input: &str) -> usize {
    let sections: Vec<&str> = input.trim().splitn(2, "\n\n").collect();
    let mut dots: HashSet<Point2d> = sections[0].lines().map(|l| {
        let vs: Vec<u32> = l.splitn(2, ",").map(|x| x.parse::<u32>().unwrap()).collect();
        return Point2d { x: vs[0], y: vs[1] };
    }).collect();
    let folds: Vec<Fold> = sections[1].lines().map(|l| {
        let end: &str = l.trim().split(" ").collect::<Vec<&str>>()[2];
        let vs: Vec<&str> = end.splitn(2, "=").collect();
        return Fold {
            axis: vs[0].chars().next().unwrap(),
            line: vs[1].parse::<u32>().unwrap(),
        };
    }).collect();

    for fold in folds {
        let axis = fold.axis;
        let line = fold.line;
        dots = dots.iter().map(|d| {
            if axis == 'y' {
                if d.y > line {
                    return Point2d {
                        x: d.x,
                        y: line - (d.y - line),
                    };
                } else {
                    return *d;
                }
            }

            if d.x > line {
                return Point2d {
                    x: line - (d.x - line),
                    y: d.y,
                };
            } else {
                return *d;
            }
        }).collect();
    }

    let w = dots.iter().map(|d| d.x).max().unwrap();
    let h = dots.iter().map(|d| d.y).max().unwrap();

    for i in 0..h+1 {
        for j in 0..w+1 {
            let c = if dots.contains(&Point2d{ x: j, y: i }) {
                '#'
            } else {
                ' '
            };
            print!("{}{}", c, c);
        }
        println!("");
    }

    return 0;
}

pub fn run() {
    let input = include_str!("../input/day13");
    println!("13.1: {:?}", part1(input));
    println!("13.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5"), 17);
    }
}
