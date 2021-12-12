use std::collections::HashMap;

fn count_paths_rewrite<'a>(
    start: &'a str,
    end: &'a str,
    neighbor_map: &'a HashMap<&'a str, Vec<&'a str>>,
    path: &mut Vec<&'a str>,
    seen: bool,
) -> u32 {
    if start == "end" {
        return 1;
    }

    let mut has_seen = seen;
    if path.contains(&start) {
        if start == "start" {
            return 0;
        }
        if start.chars().all(|c| c.is_lowercase()) {
            if has_seen {
                return 0;
            }
            has_seen = true;
        }
    }

    path.push(start);
    let mut count: u32 = 0;
    for n in neighbor_map[start].iter() {
        count += count_paths_rewrite(n, end, neighbor_map, path, has_seen);
    }
    path.pop();
    return count;
}

fn rewrite(input: &str) {
    let pairs: Vec<Vec<&str>> = input
        .trim()
        .lines()
        .map(|l| l.splitn(2, "-").collect())
        .collect();

    let mut neighbor_map: HashMap<&str, Vec<&str>> = HashMap::new();

    for pair in pairs {
        let l = pair[0];
        let r = pair[1];
        neighbor_map.entry(l).or_insert(Vec::new()).push(r);
        neighbor_map.entry(r).or_insert(Vec::new()).push(l);
    }

    println!("12.1: {:?}", count_paths_rewrite("start", "end", &neighbor_map, &mut Vec::new(), true));
    println!("12.2: {:?}", count_paths_rewrite("start", "end", &neighbor_map, &mut Vec::new(), false));
}

pub fn run() {
    let input = include_str!("../input/day12");
    rewrite(input);
}

#[cfg(test)]
mod test {
    use super::*;
    use std::collections::HashSet;

fn count_paths(
    start: String,
    end: String,
    neighbor_map: &HashMap<String, Vec<String>>,
    visited: &HashSet<String>,
) -> u32 {
    if start == "end" {
        return 1;
    }

    if start == "start" && visited.contains("start") {
        return 0;
    }

    if start.to_lowercase() == start && visited.contains(&start) {
        return 0;
    }

    let mut my_visited = visited.clone();
    my_visited.insert(start.to_string());

    let mut count: u32 = 0;
    for n in neighbor_map[&start].iter() {
        count += count_paths(
            n.to_string(),
            end.to_string(),
            neighbor_map,
            &my_visited.clone(),
        );
    }
    return count;
}

fn part1(input: &str) -> u32 {
    let pairs: Vec<Vec<&str>> = input
        .trim()
        .lines()
        .map(|l| l.splitn(2, "-").collect())
        .collect();

    let mut neighbor_map: HashMap<String, Vec<String>> = HashMap::new();

    for pair in pairs {
        let l = pair[0];
        let r = pair[1];
        neighbor_map
            .entry(l.to_string())
            .or_insert(Vec::new())
            .push(r.to_string());
        neighbor_map
            .entry(r.to_string())
            .or_insert(Vec::new())
            .push(l.to_string());
    }

    return count_paths(
        String::from("start"),
        String::from("end"),
        &neighbor_map,
        &HashSet::new(),
    );
}

fn count_paths_2<'a>(
    start: &'a str,
    end: &'a str,
    neighbor_map: &'a HashMap<&'a str, Vec<&'a str>>,
    path: &mut Vec<&'a str>,
    seen: bool,
) -> u32 {
    if start == "end" {
        return 1;
    }

    let mut has_seen = seen;
    if path.contains(&start) {
        if start == "start" {
            return 0;
        }
        if start.chars().all(|c| c.is_lowercase()) {
            if has_seen {
                return 0;
            }
            has_seen = true;
        }
    }

    path.push(start);
    let mut count: u32 = 0;
    for n in neighbor_map[start].iter() {
        count += count_paths_2(n, end, neighbor_map, path, has_seen);
    }
    path.pop();
    return count;
}

fn part2(input: &str) -> u32 {
    let pairs: Vec<Vec<&str>> = input
        .trim()
        .lines()
        .map(|l| l.splitn(2, "-").collect())
        .collect();

    let mut neighbor_map: HashMap<&str, Vec<&str>> = HashMap::new();

    for pair in pairs {
        let l = pair[0];
        let r = pair[1];
        neighbor_map.entry(l).or_insert(Vec::new()).push(r);
        neighbor_map.entry(r).or_insert(Vec::new()).push(l);
    }

    return count_paths_2("start", "end", &neighbor_map, &mut Vec::new(), false);
}

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "start-A
start-b
A-c
A-b
b-d
A-end
b-end"
            ),
            10
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "start-A
start-b
A-c
A-b
b-d
A-end
b-end"
            ),
            36
        );
    }
}
