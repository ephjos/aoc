use std::cmp::Ordering;
use std::collections::BinaryHeap;
use std::collections::{HashMap, VecDeque};

// min-heap for dijkstra
// https://doc.rust-lang.org/std/collections/binary_heap/index.html
#[derive(Copy, Clone, Eq, PartialEq)]
struct Node {
    position: (u32, u32),
    cost: u32,
}

impl Ord for Node {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .cost
            .cmp(&self.cost)
            .then_with(|| self.position.cmp(&other.position))
    }
}

// `PartialOrd` needs to be implemented as well.
impl PartialOrd for Node {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn part1(input: &str) -> u32 {
    let mut grid: HashMap<(u32, u32), u32> = HashMap::new();

    let mut max_i: u32 = 0;
    let mut max_j: u32 = 0;

    for (i, line) in input.trim().lines().enumerate() {
        for (j, c) in line.chars().into_iter().enumerate() {
            grid.insert((i as u32, j as u32), c.to_digit(10).unwrap());
            max_j = j as u32;
        }
        max_i = i as u32;
    }

    // Dijkstra
    let mut q: BinaryHeap<Node> = BinaryHeap::new();
    let mut dist: HashMap<(u32, u32), u32> = HashMap::new();
    let mut neighbors: HashMap<(u32, u32), Vec<(u32, u32)>> = HashMap::new();

    for (v, _) in grid.iter() {
        dist.insert(*v, u32::MAX);

        let mut potential_neighbors = Vec::new();
        let i = v.0;
        let j = v.1;

        if i > 0 {
            potential_neighbors.push((i - 1, j));
        }
        if i < max_i {
            potential_neighbors.push((i + 1, j));
        }
        if j > 0 {
            potential_neighbors.push((i, j - 1));
        }
        if j < max_j {
            potential_neighbors.push((i, j + 1));
        }
        neighbors.insert(*v, potential_neighbors.clone());
    }
    dist.insert((0, 0), 0);
    q.push(Node {
        position: (0, 0),
        cost: 0,
    });

    while let Some(Node { position, cost }) = q.pop() {
        let u = position;

        if u == (max_i, max_j) {
            return cost;
        }

        if cost > dist[&position] {
            continue;
        }

        for v in neighbors.get(&u).unwrap() {
            let danger = grid.get(&v).unwrap();
            let alt = if dist[&u] == u32::MAX {
                *danger
            } else {
                dist[&u] + danger
            };
            if alt < dist[&v] {
                let next = Node {
                    position: *v,
                    cost: alt,
                };
                q.push(next);
                dist.insert(*v, alt);
            }
        }
    }

    return 0;
}

fn part2(input: &str) -> u32 {
    let mut base_grid: HashMap<(u32, u32), u32> = HashMap::new();

    let mut max_i: u32 = 0;
    let mut max_j: u32 = 0;

    for (i, line) in input.trim().lines().enumerate() {
        for (j, c) in line.chars().into_iter().enumerate() {
            base_grid.insert((i as u32, j as u32), c.to_digit(10).unwrap());
            max_j = j as u32;
        }
        max_i = i as u32;
    }

    let y = max_i+1;
    let x = max_j+1;
    let mut grid: HashMap<(u32, u32), u32> = HashMap::new();

    for a in 0..5 {
        for b in 0..5 {
            let c = a*y;
            let d = b*x;
            for ((i,j), cost) in &base_grid {
                let nc = (cost+a+b) % 9;
                let nnc = if nc == 0 { 9 } else { nc };
                grid.insert((i+c,j+d), nnc);
                max_i = (i+c).max(max_i);
                max_j = (j+d).max(max_j);
            }
        }
    }

    // Dijkstra
    let mut q: BinaryHeap<Node> = BinaryHeap::new();
    let mut dist: HashMap<(u32, u32), u32> = HashMap::new();
    let mut neighbors: HashMap<(u32, u32), Vec<(u32, u32)>> = HashMap::new();

    for (v, _) in grid.iter() {
        dist.insert(*v, u32::MAX);

        let mut potential_neighbors = Vec::new();
        let i = v.0;
        let j = v.1;

        if i > 0 {
            potential_neighbors.push((i - 1, j));
        }
        if i < max_i {
            potential_neighbors.push((i + 1, j));
        }
        if j > 0 {
            potential_neighbors.push((i, j - 1));
        }
        if j < max_j {
            potential_neighbors.push((i, j + 1));
        }
        neighbors.insert(*v, potential_neighbors.clone());
    }
    dist.insert((0, 0), 0);
    q.push(Node {
        position: (0, 0),
        cost: 0,
    });

    while let Some(Node { position, cost }) = q.pop() {
        let u = position;

        if u == (max_i, max_j) {
            return cost;
        }

        if cost > dist[&position] {
            continue;
        }

        for v in neighbors.get(&u).unwrap() {
            let danger = grid.get(&v).unwrap();
            let alt = if dist[&u] == u32::MAX {
                *danger
            } else {
                dist[&u] + danger
            };
            if alt < dist[&v] {
                let next = Node {
                    position: *v,
                    cost: alt,
                };
                q.push(next);
                dist.insert(*v, alt);
            }
        }
    }

    return 0;
}

pub fn run() {
    let input = include_str!("../input/day15");
    println!("15.1: {:?}", part1(input));
    println!("15.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581"
            ),
            40
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581"
            ),
            315
        );
    }
}
