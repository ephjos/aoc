
fn part1(input: &str) -> usize {
    let lines = input.trim().lines().collect::<Vec<&str>>();
    let mut p1 = lines[0].trim().chars().last().unwrap().to_digit(10).unwrap() as usize;
    let mut p2 = lines[1].trim().chars().last().unwrap().to_digit(10).unwrap() as usize;
    let mut score1 = 0;
    let mut score2 = 0;
    let mut die = 1;
    let mut rolls = 0;

    while score1 < 1000 && score2 < 1000 {
        let mut roll1 = die;
        die = (die%100) + 1;
        roll1 += die;
        die = (die%100) + 1;
        roll1 += die;
        die = (die%100) + 1;

        p1 = ((p1+roll1-1)%10) + 1;
        score1 += p1;
        rolls += 3;
        if score1 >= 1000 {
            break;
        }

        let mut roll2 = die;
        die = (die%100) + 1;
        roll2 += die;
        die = (die%100) + 1;
        roll2 += die;
        die = (die%100) + 1;

        p2 = ((p2+roll2-1)%10) + 1;
        score2 += p2;
        rolls += 3;
        if score2 >= 1000 {
            break;
        }
    }

    let loser = score1.min(score2);
    return rolls * loser;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Game {
    p1: usize,
    p2: usize,
    score1: usize,
    score2: usize,
    p1_turn: bool,
}

impl Game {
    pub fn play(self: &mut Self) -> (usize, usize) {
        if self.score1 >= 21 {
            return (1,0);
        }
        if self.score2 >= 21 {
            return (0,1);
        }

        let mut r1 = 0;
        let mut r2 = 0;
        // While the die can return 3 numbers, each player rolls it 3 times. This
        // means we only care about the sum of three rolls. While there are
        // 27 combinations, there are only 7 sums. This array contains
        // pairs where the left number is the sum of the 3 rolls and the
        // right is the amount of times this sum happens. This allows
        // us to save on duplicate recursive calls, and simply multiply
        // the result of the games played out by the frequency of that
        // roll happening.
        let rolls = [(3,1),(4,3),(5,6),(6,7),(7,6),(8,3),(9,1)];

        if self.p1_turn {
            for (x, y) in rolls {
                let mut g = self.clone();
                g.p1 = ((g.p1 + x-1)%10)+1;
                g.score1 += g.p1;
                g.p1_turn = !g.p1_turn;
                let (a,b) = g.play();
                r1 += y * a; r2 += y * b;
            }
        } else {
            for (x, y) in rolls {
                let mut g = self.clone();
                g.p2 = ((g.p2 + x-1)%10)+1;
                g.score2 += g.p2;
                g.p1_turn = !g.p1_turn;
                let (a,b) = g.play();
                r1 += y * a; r2 += y * b;
            }
        }
        self.p1_turn = !self.p1_turn;

        return (r1, r2);
    }
}

fn part2(input: &str) -> usize {
    let lines = input.trim().lines().collect::<Vec<&str>>();
    let p1 = lines[0].trim().chars().last().unwrap().to_digit(10).unwrap() as usize;
    let p2 = lines[1].trim().chars().last().unwrap().to_digit(10).unwrap() as usize;

    let mut game = Game {
        p1,
        p2,
        score1: 0,
        score2: 0,
        p1_turn: true,
    };

    let (r1, r2) = game.play();
    return r1.max(r2);
}

pub fn run() {
    let input = include_str!("../input/day21");
    println!("21.1: {:?}", part1(input));
    println!("21.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("Player 1 starting position: 4
Player 2 starting position: 8"), 739785);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("Player 1 starting position: 4
Player 2 starting position: 8"), 444356092776315);
    }
}

