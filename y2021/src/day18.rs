
fn part1(input: &str) -> isize {
    return 0;
}

fn part2(input: &str) -> isize {
    return 0;
}

pub fn run() {
    let input = include_str!("../input/day18");
    println!("18.1: {:?}", part1(input));
    println!("18.2: {:?}", part2(input));
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

