#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone, Copy)]
struct Projectile {
    px: i64,
    py: i64,
    vx: i64,
    vy: i64,
}

impl Projectile {
    fn step(self: &mut Self) {
        self.px += self.vx;
        self.py += self.vy;
        self.vx -= self.vx.signum();
        self.vy -= 1;
    }
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone, Copy)]
struct Target {
    min_x: i64,
    max_x: i64,
    min_y: i64,
    max_y: i64,
    width: i64,
    height: i64,
}

impl Target {
    fn contains_projectile(self: &Self, p: &Projectile) -> bool {
        return (self.min_x <= p.px && p.px <= self.max_x)
            && (self.min_y <= p.py && p.py <= self.max_y);
    }

    fn project_missed(self: &Self, p: &Projectile) -> bool {
        return (p.py < self.min_y) || (p.px > self.max_x);
    }
}

fn part1(input: &str) -> i64 {
    let nums = &input.trim()[13..];
    let parts = nums
        .splitn(2, ", ")
        .map(|s| {
            (&s[2..])
                .splitn(2, "..")
                .map(|x| x.parse::<i64>().unwrap())
                .collect::<Vec<i64>>()
        })
        .collect::<Vec<Vec<i64>>>();
    let min_x = *parts[0].iter().min().unwrap();
    let max_x = *parts[0].iter().max().unwrap();
    let min_y = *parts[1].iter().min().unwrap();
    let max_y = *parts[1].iter().max().unwrap();
    let target = Target {
        min_x,
        max_x,
        min_y,
        max_y,
        width: max_y - min_y,
        height: max_x - min_x,
    };


    // Precompute min_vx using gaussian sum
    let mut min_vx = 0;
    while min_vx * (min_vx+1)/2 < target.min_x {
        min_vx += 1;
    }

    let max_vx = target.max_x+1;
    let min_vy = target.min_y-1;

    let mut max_y = i64::MIN;
    for vx in min_vx..max_vx {
        for vy in min_vy..min_vy.abs() {
            let p = &mut Projectile {
                px: 0,
                py: 0,
                vx,
                vy,
            };

            let mut this_max_y = i64::MIN;
            let mut hit = false;
            while !target.project_missed(p) {
                this_max_y = p.py.max(this_max_y);
                if target.contains_projectile(p) {
                    hit = true;
                    break;
                }

                p.step();
            }

            if hit {
                if this_max_y > max_y {
                    max_y = this_max_y;
                }
            }
        }
    }

    return max_y;
}

fn part2(input: &str) -> i64 {
    let nums = &input.trim()[13..];
    let parts = nums
        .splitn(2, ", ")
        .map(|s| {
            (&s[2..])
                .splitn(2, "..")
                .map(|x| x.parse::<i64>().unwrap())
                .collect::<Vec<i64>>()
        })
        .collect::<Vec<Vec<i64>>>();
    let min_x = *parts[0].iter().min().unwrap();
    let max_x = *parts[0].iter().max().unwrap();
    let min_y = *parts[1].iter().min().unwrap();
    let max_y = *parts[1].iter().max().unwrap();
    let target = Target {
        min_x,
        max_x,
        min_y,
        max_y,
        width: max_y - min_y,
        height: max_x - min_x,
    };


    // Precompute min_vx using gaussian sum
    let mut min_vx = 0;
    while min_vx * (min_vx+1)/2 < target.min_x {
        min_vx += 1;
    }

    let max_vx = target.max_x+1;
    let min_vy = target.min_y-1;

    let mut count = 0;
    for vx in min_vx..max_vx {
        for vy in min_vy..min_vy.abs() {
            let p = &mut Projectile {
                px: 0,
                py: 0,
                vx,
                vy,
            };

            let mut hit = false;
            while !target.project_missed(p) {
                if target.contains_projectile(p) {
                    hit = true;
                    break;
                }

                p.step();
            }

            if hit {
                count += 1;
            }
        }
    }

    return count;
}

pub fn run() {
    let input = include_str!("../input/day17");
    println!("17.1: {:?}", part1(input));
    println!("17.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("target area: x=20..30, y=-10..-5"), 45);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("target area: x=20..30, y=-10..-5"), 112);
    }
}
