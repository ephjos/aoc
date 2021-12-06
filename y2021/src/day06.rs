
use super::com::*;

fn part1(input: &str) -> usize {
    fn fish_list_parser<'a>() -> impl Parser<'a, Vec<usize>> {
        one_or_more(left(uinteger, whitespace_wrap(zero_or_more(match_literal(",")))))
    }

    let fish_list = fish_list_parser().parse(input).unwrap().1;

    let mut fish_ages = vec![0usize; 7];
    for fish in fish_list {
        fish_ages[fish] += 1;
    }
    let mut new_fish_ages = vec![0usize; 9];

    for _ in 0..80 {
        let temp = fish_ages[0];
        fish_ages[0] = fish_ages[1];
        fish_ages[1] = fish_ages[2];
        fish_ages[2] = fish_ages[3];
        fish_ages[3] = fish_ages[4];
        fish_ages[4] = fish_ages[5];
        fish_ages[5] = fish_ages[6];
        fish_ages[6] = temp;

        let temp2 = new_fish_ages[0];
        new_fish_ages[0] = new_fish_ages[1];
        new_fish_ages[1] = new_fish_ages[2];
        new_fish_ages[2] = new_fish_ages[3];
        new_fish_ages[3] = new_fish_ages[4];
        new_fish_ages[4] = new_fish_ages[5];
        new_fish_ages[5] = new_fish_ages[6];
        new_fish_ages[6] = new_fish_ages[7];
        new_fish_ages[7] = new_fish_ages[8];
        new_fish_ages[8] = temp+temp2;

        fish_ages[6] += temp2;
    }

    return fish_ages.iter().sum::<usize>() + new_fish_ages.iter().sum::<usize>();
}

fn part2(input: &str) -> usize {
    fn fish_list_parser<'a>() -> impl Parser<'a, Vec<usize>> {
        one_or_more(left(uinteger, whitespace_wrap(zero_or_more(match_literal(",")))))
    }

    let fish_list = fish_list_parser().parse(input).unwrap().1;

    let mut fish_ages = vec![0usize; 7];
    for fish in fish_list {
        fish_ages[fish] += 1;
    }
    let mut new_fish_ages = vec![0usize; 9];

    for _ in 0..256 {
        let temp = fish_ages[0];
        fish_ages[0] = fish_ages[1];
        fish_ages[1] = fish_ages[2];
        fish_ages[2] = fish_ages[3];
        fish_ages[3] = fish_ages[4];
        fish_ages[4] = fish_ages[5];
        fish_ages[5] = fish_ages[6];
        fish_ages[6] = temp;

        let temp2 = new_fish_ages[0];
        new_fish_ages[0] = new_fish_ages[1];
        new_fish_ages[1] = new_fish_ages[2];
        new_fish_ages[2] = new_fish_ages[3];
        new_fish_ages[3] = new_fish_ages[4];
        new_fish_ages[4] = new_fish_ages[5];
        new_fish_ages[5] = new_fish_ages[6];
        new_fish_ages[6] = new_fish_ages[7];
        new_fish_ages[7] = new_fish_ages[8];
        new_fish_ages[8] = temp+temp2;

        fish_ages[6] += temp2;
    }

    return fish_ages.iter().sum::<usize>() + new_fish_ages.iter().sum::<usize>();
}

pub fn run() {
    let input = include_str!("../input/day06");
    println!("06.1: {:?}", part1(input));
    println!("06.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("3,4,3,1,2"), 5934);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("3,4,3,1,2"), 26984457539);
    }
}
