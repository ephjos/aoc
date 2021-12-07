
fn part1(input: &str) -> isize {
    let mut crabs = input.trim().split(",")
        .map(|x| x.parse::<isize>().unwrap())
        .collect::<Vec<_>>();

    crabs.sort();

    let med = crabs[crabs.len() / 2];
    let cost = crabs.iter().map(|x| (x-med).abs()).sum::<isize>();

    return cost;
}

fn part2(input: &str) -> isize {
    let crabs = input.trim().split(",")
        .map(|x| x.parse::<isize>().unwrap())
        .collect::<Vec<_>>();

    let mean = crabs.iter().sum::<isize>() as f64 / crabs.len() as f64;
    let mean_floor = mean.floor() as isize;
    let mean_ceil = mean.ceil() as isize;

    fn gauss(n: isize) -> isize {
        return n * (n+1) / 2;
    }

    let floor_cost = crabs.iter().map(|x| gauss((x-mean_floor).abs())).sum::<isize>();
    let ceil_cost = crabs.iter().map(|x| gauss((x-mean_ceil).abs())).sum::<isize>();

    if floor_cost < ceil_cost {
        return floor_cost;
    }
    return ceil_cost;
}

pub fn run() {
    let input = include_str!("../input/day07");
    println!("07.1: {:?}", part1(input));
    println!("07.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("16,1,2,0,4,2,7,1,2,14"), 37)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("16,1,2,0,4,2,7,1,2,14"), 168)
    }
}
