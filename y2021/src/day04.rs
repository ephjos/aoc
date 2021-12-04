
#[derive(Debug, Clone)]
struct Board {
    grid: Vec<Vec<i32>>,
}

impl Board {
    fn mark_and_check(&mut self, draw: i32) -> i32 {
        for i in 0..self.grid.len() {
            for j in 0..self.grid[i].len() {
                if self.grid[i][j] == draw {
                    self.grid[i][j] = -1;
                }
            }
        }

        // Rows
        for i in 0..self.grid.len() {
            let mut res = true;
            for j in 0..self.grid[i].len() {
                res &= self.grid[i][j] == -1;
            }
            if res {
                return draw * self.score();
            }
        }

        // Cols
        for i in 0..self.grid.len() {
            let mut res = true;
            for j in 0..self.grid[i].len() {
                res &= self.grid[j][i] == -1;
            }
            if res {
                return draw * self.score();
            }
        }

        // Diag TL -> BR
        let mut res = false;
        for i in 0..self.grid.len() {
            res &= self.grid[i][i] == -1;
        }
        if res {
            return draw * self.score();
        }

        // Diag BL -> TR
        let mut res = false;
        for i in 0..self.grid.len() {
            res &= self.grid[i][self.grid.len()-1-i] == -1;
        }
        if res {
            return draw * self.score();
        }

        return -1;
    }

    fn score(&self) -> i32 {
        let mut res: i32 = 0;

        for row in &self.grid {
            for value in row {
                if *value != -1 {
                    res += value;
                }
            }
        }
        return res;
    }
}

fn part1(input: &str) -> i32 {
    let lines = input.lines().collect::<Vec<_>>();
    let draws = lines[0].split(",").map(|x| x.parse::<i32>().unwrap()).collect::<Vec<_>>();

    let mut i = 2;
    let mut boards: Vec<Board> = Vec::new();
    while i < lines.len() {
        let mut board = Board {
            grid: vec![vec![0 as i32; 5]; 5],
        };

        for k in 0..5 {
            for (j, v) in lines[i].split_whitespace().map(|x| x.parse::<i32>().unwrap()).enumerate() {
                board.grid[k][j] = v;
            }

            i += 1;
        }

        i+=1;
        boards.push(board);
    }

    for draw in draws {
        for board in &mut boards {
            let res = board.mark_and_check(draw);
            if res != -1 {
                return res;
            }
        }
    }

    return 0;
}

fn part2(input: &str) -> i32 {
    let lines = input.lines().collect::<Vec<_>>();
    let draws = lines[0].split(",").map(|x| x.parse::<i32>().unwrap()).collect::<Vec<_>>();

    let mut i = 2;
    let mut boards: Vec<Board> = Vec::new();
    while i < lines.len() {
        let mut board = Board {
            grid: vec![vec![0 as i32; 5]; 5],
        };

        for k in 0..5 {
            for (j, v) in lines[i].split_whitespace().map(|x| x.parse::<i32>().unwrap()).enumerate() {
                board.grid[k][j] = v;
            }

            i += 1;
        }

        i+=1;
        boards.push(board);
    }

    for draw in draws {
        let mut i = 0;
        while i < boards.len() {
            let res = boards[i].mark_and_check(draw);
            if res != -1 {
                if boards.len() == 1 {
                    return res;
                }
                boards.remove(i);
                continue;
            }
            i += 1;
        }
    }

    return 0;
}

pub fn run() {
    let input = include_str!("../input/day04");
    println!("04.1: {:?}", part1(input));
    println!("04.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7"), 4512);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7"), 1924);
    }
}
