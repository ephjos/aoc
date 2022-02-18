use std::collections::HashMap;

#[derive(Debug, Hash, PartialEq, Eq, PartialOrd, Ord, Clone, Copy)]
struct Point2D {
    x: isize,
    y: isize,
}

impl Point2D {
    pub fn new() -> Point2D {
        Point2D {
            x: 0,
            y: 0,
        }
    }
}

#[derive(Debug, PartialEq, Eq)]
struct Message {
    algorithm: Vec<bool>,
    image: HashMap<Point2D, bool>,
    space: bool,
    min: Point2D,
    max: Point2D,
}

impl Message {
    pub fn parse(input: &str) -> Message {
        let blocks = input.trim().splitn(2, "\n\n").collect::<Vec<&str>>();

        let algorithm = blocks[0].trim().chars().map(|c| match c {
            '.' => false,
            '#' => true,
            x => panic!("Unknown algorithm char '{}'", x),
        }).collect::<Vec<bool>>();

        let min = Point2D::new();
        let mut max = Point2D::new();
        let mut image = HashMap::new();
        let lines = &blocks[1].lines().collect::<Vec<&str>>();
        for i in 0..lines.len() {
            for j in 0..lines[i].len() {
                let c = lines[i].chars().nth(j).unwrap();
                let x = j as isize;
                let y = i as isize;
                image.insert(Point2D { x, y }, match c {
                    '.' => false,
                    '#' => true,
                    u => panic!("Unknown image char '{}'", u),
                });
                max.x = x;
                max.y = y;
            }
        }

        let space = algorithm[0];

        Message {
            algorithm,
            image,
            min,
            max,
            space,
        }
    }

    pub fn print_image(self: &Self) {
        let o = 2;
        let min_x = self.min.x-o;
        let min_y = self.min.y-o;
        let max_x = self.max.x+o+1;
        let max_y = self.max.y+o+1;
        for i in min_y..max_y {
            for j in min_x..max_x {
                let x = j as isize;
                let y = i as isize;
                let key = Point2D { x, y };
                if let Some(b) = self.image.get(&key) {
                    if *b {
                        print!("#");
                        continue;
                    }
                }
                print!(".");
            }
            println!();
        }
    }

    pub fn enhance(self: &mut Self) {
        let min_x = self.min.x-1;
        let min_y = self.min.y-1;
        let max_x = self.max.x+2;
        let max_y = self.max.y+2;

        let mut new_image = HashMap::new();

        self.space = self.algorithm[(self.space as usize) * 511];

        for i in min_y..max_y {
            for j in min_x..max_x {
                let jj = j as isize;
                let ii = i as isize;

                let mut buf = 0;
                for u in -1..=1 {
                    for v in -1..=1 {
                        let x = jj + v;
                        let y = ii + u;
                        let key = Point2D { x, y };
                        if let Some(b) = self.image.get(&key) {
                            buf += *b as usize;
                        } else {
                            buf += self.space as usize;
                        }
                        buf <<= 1;
                    }
                }
                buf >>= 1;
                let key = Point2D { x: jj, y: ii };
                new_image.insert(key, self.algorithm[buf]);
            }
        }

        self.image.clear();
        self.image = new_image;
        self.min.x = min_x;
        self.min.y = min_y;
        self.max.x = max_x;
        self.max.y = max_y;
    }

    pub fn count(self: &Self) -> usize {
        return self.image.values().filter(|b| **b).count();
    }
}

fn part1(input: &str) -> usize {
    let mut message = Message::parse(input);
    message.enhance();
    message.enhance();
    return message.count();
}

fn part2(input: &str) -> usize {
    let mut message = Message::parse(input);
    for _ in 0..50 {
        message.enhance();
    }
    message.print_image();
    return message.count();
}

pub fn run() {
    let input = include_str!("../input/day20");
    println!("20.1: {:?}", part1(input));
    println!("20.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###"), 35);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###"), 3351);
    }
}

