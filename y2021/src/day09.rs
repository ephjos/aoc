use std::collections::{HashSet, VecDeque, BinaryHeap};

fn part1(input: &str) -> u32 {
    let grid: Vec<Vec<u32>> = input
        .lines()
        .map(|line|
            line.chars().map(|c| c.to_digit(10).unwrap()).collect::<Vec<u32>>())
        .collect();

    let mut min_indices: Vec<(usize, usize)> = Vec::new();

    let h = grid.len();
    let w = grid[0].len();
    for i in 0..h {
        for j in 0..w {
            let curr = grid[i][j];
            let has_up = i > 0;
            let has_down = i < h-1;
            let has_left = j > 0;
            let has_right = j < w-1;

            let mut is_min = true;
            if has_up {
                is_min &= curr < grid[i-1][j];
            }
            if has_down {
                is_min &= curr < grid[i+1][j];
            }
            if has_left {
                is_min &= curr < grid[i][j-1];
            }
            if has_right {
                is_min &= curr < grid[i][j+1];
            }

            if is_min {
                min_indices.push((i,j));
            }
        }
    }

    return min_indices.iter().map(|(i,j)| grid[*i][*j]+1).sum::<u32>();
}

fn bfs_basin(grid: &Vec<Vec<u32>>, p: &(usize, usize)) -> u32 {
    let h = grid.len();
    let w = grid[0].len();
    let mut q: VecDeque<(usize, usize)> = VecDeque::new();
    let mut explored = HashSet::new();
    explored.insert(*p);
    q.push_back(*p);

    while !q.is_empty() {
        let (i, j) = q.pop_front().unwrap();
        let mut neighbors = Vec::new();

        if i > 0 && grid[i-1][j] != 9 {
            neighbors.push((i-1, j));
        }

        if i < h-1 && grid[i+1][j] != 9 {
            neighbors.push((i+1, j));
        }

        if j > 0 && grid[i][j-1] != 9 {
            neighbors.push((i, j-1));
        }

        if j < w-1 && grid[i][j+1] != 9 {
            neighbors.push((i, j+1));
        }

        for n in neighbors {
            if !explored.contains(&n) {
                explored.insert(n);
                q.push_back(n);
            }
        }
    }

    return explored.len() as u32;
}

fn part2(input: &str) -> u32 {
    let grid: Vec<Vec<u32>> = input
        .lines()
        .map(|line|
            line.chars().map(|c| c.to_digit(10).unwrap()).collect::<Vec<u32>>())
        .collect();

    let mut min_indices: Vec<(usize, usize)> = Vec::new();

    let h = grid.len();
    let w = grid[0].len();
    for i in 0..h {
        for j in 0..w {
            let curr = grid[i][j];
            let has_up = i > 0;
            let has_down = i < h-1;
            let has_left = j > 0;
            let has_right = j < w-1;

            let mut is_min = true;
            if has_up {
                is_min &= curr < grid[i-1][j];
            }
            if has_down {
                is_min &= curr < grid[i+1][j];
            }
            if has_left {
                is_min &= curr < grid[i][j-1];
            }
            if has_right {
                is_min &= curr < grid[i][j+1];
            }

            if is_min {
                min_indices.push((i,j));
            }
        }
    }

    let mut basins = min_indices.iter().map(|p| bfs_basin(&grid, p)).collect::<BinaryHeap<u32>>();

    return basins.pop().unwrap() * basins.pop().unwrap() * basins.pop().unwrap();
}

pub fn run() {
    let input = include_str!("../input/day09");
    println!("09.1: {:?}", part1(input));
    println!("09.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "2199943210
3987894921
9856789892
8767896789
9899965678"
            ),
            15
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "2199943210
3987894921
9856789892
8767896789
9899965678"
            ),
            1134
        );
    }
}
