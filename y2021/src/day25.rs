use std::collections::HashSet;

const EAST: char = '>';
const SOUTH: char = 'v';

struct State {
    w: usize,
    h: usize,
    east: HashSet<(usize, usize)>,
    south: HashSet<(usize, usize)>,
}

impl std::fmt::Debug for State {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut out = String::new();
        for i in 0..self.h {
            for j in 0..self.w {
                if self.east.contains(&(i,j)) {
                    out.push('>');
                } else if self.south.contains(&(i,j)) {
                    out.push('v');
                } else {
                    out.push('.');
                }
            }
            out.push('\n');
        }
        write!(f, "{}", out)
    }
}

impl State {
    fn parse(input: &str) -> State {
        let mut w = 0;
        let mut h = 0;
        let mut east = HashSet::new();
        let mut south = HashSet::new();
        let lines = input.trim().lines().collect::<Vec<&str>>();

        for i in 0..lines.len() {
            let line = lines[i].chars().collect::<Vec<char>>();
            for j in 0..line.len() {
                let c = line[j];
                if c == EAST {
                    east.insert((i,j));
                } else if c == SOUTH {
                    south.insert((i,j));
                }
                w = j;
            }
            h = i;
        }

        w += 1;
        h += 1;

        State {
            w,
            h,
            east,
            south,
        }
    }

    fn get_move(self: &Self, src: (usize, usize)) -> Option<(usize, usize)> {
        let dst = if self.east.get(&src).is_some() {
            // check right
            (src.0, (src.1 + 1) % self.w)
        } else {
            // check below
            ((src.0 + 1) % self.h, src.1)
        };

        if self.east.contains(&dst) || self.south.contains(&dst) {
            return None;
        }

        return Some(dst);
    }

    fn all_stopped(self: &Self) -> bool {
        let all_stopped_east = self.east.iter().all(|p| self.get_move(*p).is_none());
        let all_stopped_south = self.south.iter().all(|p| self.get_move(*p).is_none());
        return all_stopped_east && all_stopped_south;
    }

    fn step(self: &mut Self) {
        let mut new_east: HashSet<(usize, usize)> = HashSet::new();

        for i in 0..self.h {
            for j in 0..self.w {
                if self.east.contains(&(i,j)) {
                    if let Some(x) = self.get_move((i,j)) {
                        new_east.insert(x);
                    } else {
                        new_east.insert((i,j));
                    }
                }
            }
        }

        self.east = new_east;
        let mut new_south: HashSet<(usize, usize)> = HashSet::new();

        for i in 0..self.h {
            for j in 0..self.w {
                if self.south.contains(&(i,j)) {
                    if let Some(x) = self.get_move((i,j)) {
                        new_south.insert(x);
                    } else {
                        new_south.insert((i,j));
                    }
                }
            }
        }
        self.south = new_south;
    }
}

fn part1(input: &str) -> isize {
    let mut state = State::parse(input);
    let mut count = 1;

    while !state.all_stopped() {
        state.step();
        count += 1;
    }

    return count;
}

pub fn run() {
    let input = include_str!("../input/day25");
    println!("25: {:?}", part1(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>"), 58);
    }
}

