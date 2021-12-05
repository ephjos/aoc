use std::collections::HashMap;
use std::hash::Hash;
use std::ops;

#[derive(Debug, Hash, PartialEq, Eq, PartialOrd, Ord)]
struct Point2d {
    x: isize,
    y: isize,
}

impl Point2d {
    fn from_str(s: &str) -> Point2d {
        let vs = s.split(",").map(|x| x.parse::<isize>().unwrap()).collect::<Vec<_>>();
        return Point2d { x: vs[0], y: vs[1] };
    }
}

impl ops::Sub for &Point2d {
    type Output = Point2d;

    fn sub(self, rhs: Self) -> Self::Output {
        return Point2d{ x: self.x - rhs.x, y: self.y - rhs.y };
    }
}


fn part1(input: &str) -> usize {
    let mut seen_points: HashMap<Point2d, usize> = HashMap::new();

    for line in input.lines() {
        let points = line.split(" -> ").map(|s| Point2d::from_str(s)).collect::<Vec<_>>();
        let start = &points[0]; let end = &points[1];
        if !((start.x == end.x) || (start.y == end.y)) {
            continue;
        }
        let diff = end - start;
        if diff.x != 0 {
            let dir = diff.x.signum();
            let y = start.y;

            let mut i = start.x;
            while i != end.x+dir {
                let key = Point2d { x: i, y };
                *seen_points.entry(key).or_insert(0) += 1;
                i += dir;
            }
        } else if diff.y != 0 {
            let dir = diff.y.signum();
            let x = start.x;

            let mut i = start.y;
            while i != end.y+dir {
                let key = Point2d { x, y: i };
                *seen_points.entry(key).or_insert(0) += 1;
                i += dir;
            }
        }
    }

    return seen_points.values().filter(|x| x > &&1).count()
}

fn part2(input: &str) -> usize {
    let mut seen_points: HashMap<Point2d, usize> = HashMap::new();

    for line in input.lines() {
        let points = line.split(" -> ").map(|s| Point2d::from_str(s)).collect::<Vec<_>>();
        let start = &points[0]; let end = &points[1];
        let diff = end - start;

        if !((start.x == end.x) || (start.y == end.y)) {
            let x_step = diff.x.signum();
            let y_step = diff.y.signum();

            let mut x = start.x;
            let mut y = start.y;

            while (x != end.x+x_step) && (y != end.y+y_step) {
                let key = Point2d { x, y };
                *seen_points.entry(key).or_insert(0) += 1;

                x += x_step;
                y += y_step;
            }
        } else if diff.x != 0 {
            let dir = diff.x.signum();
            let y = start.y;

            let mut i = start.x;
            while i != end.x+dir {
                let key = Point2d { x: i, y };
                *seen_points.entry(key).or_insert(0) += 1;
                i += dir;
            }
        } else if diff.y != 0 {
            let dir = diff.y.signum();
            let x = start.x;

            let mut i = start.y;
            while i != end.y+dir {
                let key = Point2d { x, y: i };
                *seen_points.entry(key).or_insert(0) += 1;
                i += dir;
            }
        }
    }

    return seen_points.values().filter(|x| x > &&1).count()
}

pub fn run() {
    let input = include_str!("../input/day05");
    println!("05.1: {:?}", part1(input));
    println!("05.2: {:?}", part2(input));
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"), 5);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"), 12);
    }
}
