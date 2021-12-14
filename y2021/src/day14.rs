use std::collections::HashMap;

fn part1(input: &str) -> u32 {
    let sections: Vec<&str> = input.trim().splitn(2, "\n\n").collect();
    let template: Vec<char> = sections[0].trim().chars().collect();
    let rules: HashMap<[char; 2], char> = sections[1]
        .trim()
        .lines()
        .map(|l| {
            let sides: Vec<&str> = l.splitn(2, " -> ").collect();
            let left: Vec<char> = sides[0].trim().chars().collect();
            let right: char = sides[1].trim().chars().nth(0).unwrap();
            return ([left[0], left[1]], right);
        })
        .collect();

    let mut counts: HashMap<char, u32> = HashMap::new();
    for c in &template {
        *counts.entry(*c).or_insert(0) += 1;
    }

    let mut old_template = template.clone();
    for _ in 0..10 {
        let windows = old_template.windows(2);
        let mut new_template: Vec<char> = Vec::new();
        for window in windows {
            new_template.push(window[0]);
            let new = rules[window];
            new_template.push(new);
            *counts.entry(new).or_insert(0) += 1;
        }
        new_template.push(*old_template.last().unwrap());
        old_template = new_template;
    }

    let max = counts.values().max().unwrap();
    let min = counts.values().min().unwrap();

    return max - min;
}

fn part2(input: &str) -> usize {
    let sections: Vec<&str> = input.trim().splitn(2, "\n\n").collect();
    let template: Vec<char> = sections[0].trim().chars().collect();
    let rules: HashMap<[char; 2], char> = sections[1]
        .trim()
        .lines()
        .map(|l| {
            let sides: Vec<&str> = l.splitn(2, " -> ").collect();
            let left: Vec<char> = sides[0].trim().chars().collect();
            let right: char = sides[1].trim().chars().nth(0).unwrap();
            return ([left[0], left[1]], right);
        })
        .collect();

    let mut counts: HashMap<[char; 2], usize> = HashMap::new();
    let mut indiv_counts: HashMap<char, usize> = HashMap::new();
    for (p, _) in &rules {
        counts.insert(*p, 0);
    }

    for window in template.windows(2) {
        *counts.entry([window[0], window[1]]).or_insert(0) += 1;
    }
    for t in template {
        *indiv_counts.entry(t).or_insert(0) += 1;
    }

    for _ in 0..40 {
        let old_counts = &counts.clone();
        for (p, m) in &rules {
            let c = old_counts[p];
            *counts.entry([p[0], *m]).or_insert(0) += c;
            *counts.entry([*m, p[1]]).or_insert(0) += c;
            *counts.entry(*p).or_insert(0) -= c;
            *indiv_counts.entry(*m).or_insert(0) += c;
        }
    }

    let max = indiv_counts.values().max().unwrap();
    let min = indiv_counts.values().min().unwrap();

    return max - min;
}

pub fn run() {
    let input = include_str!("../input/day14");
    println!("14.1: {:?}", part1(input));
    println!("14.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"
            ),
            1588
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"
            ),
            2188189693529
        );
    }
}
