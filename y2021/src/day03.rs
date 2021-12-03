
fn part1(input: &str) -> isize {
    let mut gamma = 0;
    let mut epsilon = 0;
    let lines = input.lines().collect::<Vec<_>>();
    let mut counts: Vec<isize> = vec![0; lines[0].len()];
    for line in lines {
        for (i, c) in line.chars().enumerate() {
            match c {
                '1' => counts[i] += 1,
                '0' => counts[i] += -1,
                _ => unreachable!(),
            }
        }
    }

    for count in counts {
        if count > 1 {
            gamma |= 1; gamma <<= 1;
            epsilon <<= 1;
        } else {
            epsilon |= 1; epsilon <<= 1;
            gamma <<= 1;
        }
    }
    gamma >>= 1;
    epsilon >>= 1;
    return gamma * epsilon;
}

fn part2(input: &str) -> isize {
    let lines = input.lines().map(|line| line.chars().collect::<Vec<char>>()).collect::<Vec<_>>();
    let mut o2_lines = lines.clone();

    let mut i = 0;
    while o2_lines.len() != 1 {
        let len = o2_lines.len();
        let zero_count: usize = o2_lines.iter().map(|line| (line[i] == '0') as usize).sum();
        let one_count = len - zero_count;
        let target = if one_count >= zero_count {
            '1'
        } else {
            '0'
        };
        o2_lines.retain(|line| line[i] == target);
        i += 1;
    }
    let o2 = isize::from_str_radix(&String::from_iter(&o2_lines[0]), 2).unwrap();

    let mut co2_lines = lines.clone();

    let mut i = 0;
    while co2_lines.len() != 1 {
        let len = co2_lines.len();
        let zero_count: usize = co2_lines.iter().map(|line| (line[i] == '0') as usize).sum();
        let one_count = len - zero_count;
        let target = if one_count >= zero_count {
            '0'
        } else {
            '1'
        };
        co2_lines.retain(|line| line[i] == target);
        i += 1;
    }
    let co2 = isize::from_str_radix(&String::from_iter(&co2_lines[0]), 2).unwrap();

    return o2 * co2;
}

pub fn run() {
    let input = include_str!("../input/day03");
    println!("03.1: {:?}", part1(input));
    println!("03.2: {:?}", part2(input));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010"), 198);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010"), 230);
    }
}
