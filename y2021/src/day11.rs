
fn step(grid: &mut Vec<Vec<u8>>) -> usize {
    let h = grid.len() as isize;
    let w = grid[0].len() as isize;

    // Inc all by 1
    for i in 0..h {
        for j in 0..w {
            grid[i as usize][j as usize] += 1;
        }
    }

    let mut flash_count = 0;
    let mut is_flash = true;
    while is_flash {
        is_flash = false;
        for i in 0..h {
            for j in 0..w {
                if grid[i as usize][j as usize] > 9 {
                    is_flash = true;
                    flash_count += 1;
                    grid[i as usize][j as usize] = 0;
                    for ii in -1..=1 {
                        for jj in -1..=1 {
                            if ii == 0 && jj == 0 {
                                continue;
                            }
                            let y = i+ii;
                            let x = j+jj;
                            if y >= 0 && y <= h-1 && x >= 0 && x <= w-1 && grid[y as usize][x as usize] != 0 {
                                grid[y as usize][x as usize] += 1;
                            }
                        }
                    }
                }
            }
        }
    }

    return flash_count;
}

fn part1(input: &str) -> usize {
    let mut grid: Vec<Vec<u8>> = Vec::new();
    for line in input.trim().lines() {
        grid.push(line.chars().map(|c| c.to_digit(10).unwrap() as u8).collect());
    }

    let mut count = 0;
    for _ in 0..100 {
        count += step(&mut grid);
    }
    return count;
}

fn part2(input: &str) -> isize {
    let mut grid: Vec<Vec<u8>> = Vec::new();
    for line in input.trim().lines() {
        grid.push(line.chars().map(|c| c.to_digit(10).unwrap() as u8).collect());
    }

    let h = grid.len();
    let w = grid[0].len();
    let mut count = 0;
    let mut all_flash = false;

    while !all_flash {
        step(&mut grid);
        all_flash = true;
        for i in 0..h {
            for j in 0..w {
                if grid[i][j] != 0 {
                    all_flash = false;
                    break;
                }
            }
            if !all_flash {
                break;
            }
        }
        count += 1;
    }
    return count;
}

pub fn run() {
    let input = include_str!("../input/day11");
    println!("11.1: {:?}", part1(input));
    println!("11.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"), 1656);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"), 195);
    }
}
