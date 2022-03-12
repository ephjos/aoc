use std::collections::HashSet;

#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
enum Tile {
    Empty,
    A,
    B,
    C,
    D,
}

impl Tile {
    fn cost(self: &Self) -> usize {
        match self {
            &Tile::A => 1,
            &Tile::B => 10,
            &Tile::C => 100,
            &Tile::D => 1000,
            &Tile::Empty => 0,
        }
    }

    fn to_char(self: &Self) -> char {
        match self {
            &Tile::A => 'A',
            &Tile::B => 'B',
            &Tile::C => 'C',
            &Tile::D => 'D',
            &Tile::Empty => '.',
        }
    }

    fn from_char(c: char) -> Tile {
        match c {
            '.' => Tile::Empty,
            'A' => Tile::A,
            'B' => Tile::B,
            'C' => Tile::C,
            'D' => Tile::D,
            _ => Tile::Empty,
        }
    }

    fn room_index(self: &Self) -> usize {
        match self {
            &Tile::A => 0,
            &Tile::B => 1,
            &Tile::C => 2,
            &Tile::D => 3,
            _ => panic!("Impossible!"),
        }
    }
}

#[derive(Clone, Hash, PartialEq, Eq)]
struct State {
    cost: usize,
    hallway: [Tile; 11],
    rooms: [Vec<Tile>; 4],
    rows: usize,
}

impl std::fmt::Debug for State {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut out = String::new();
        out.push('\n');
        out.push_str(&format!("cost: {}\n", self.cost));
        for i in 0..11 {
            out.push_str(&format!("{}", self.hallway[i].to_char()));
        }
        for i in 0..self.rows {
            out.push('\n');
            out.push_str("  ");
            for j in 0..4 {
                out.push_str(&format!("{} ", self.rooms[j][self.rows-i-1].to_char()));
            }
        }
        out.push('\n');
        return write!(f, "{}", out);
    }
}

impl State {
    fn parse(input: &str) -> State {
        let lines = input.trim().lines().collect::<Vec<&str>>();
        let hallway = [Tile::Empty; 11];
        let mut room0 = Vec::new();
        let mut room1 = Vec::new();
        let mut room2 = Vec::new();
        let mut room3 = Vec::new();
        let mut rows = 0;

        for line in &lines[2..] {
            let chars = line.chars().collect::<Vec<char>>();
            if chars[3] == '#' {
                break;
            }
            room0.insert(0, Tile::from_char(chars[3]));
            room1.insert(0, Tile::from_char(chars[5]));
            room2.insert(0, Tile::from_char(chars[7]));
            room3.insert(0, Tile::from_char(chars[9]));
            rows += 1;
        }

        State {
            cost: 0,
            hallway,
            rooms: [
                room0,
                room1,
                room2,
                room3,
            ],
            rows,
        }
    }

    fn step(self: &Self, seen: &mut HashSet<State>) {
        if self.is_finished() {
            return;
        }

        for i in 0..self.hallway.len() {
            let tile = self.hallway[i];
            if tile != Tile::Empty {
                // Letter in hallway
                let room_index = tile.room_index();
                let room_hallway_index = (room_index + 1) * 2;

                let mut hall_clear = true;
                let (l,h) = if i < room_hallway_index {
                    (i+1, room_hallway_index)
                } else {
                    (room_hallway_index, i-1)
                };

                for j in l..=h {
                    if self.hallway[j] != Tile::Empty {
                        hall_clear = false;
                        break;
                    }
                }

                if !hall_clear {
                    continue;
                }

                let room = &self.rooms[room_index];
                let can_enter = room.iter().all(|t| t == &tile || t == &Tile::Empty);

                if hall_clear && can_enter {
                    let enter_index = room.iter().position(|t| t == &Tile::Empty).unwrap();
                    let mut new_state = self.clone();
                    new_state.rooms[room_index][enter_index] = tile;
                    new_state.hallway[i] = Tile::Empty;
                    let enter_cost = room.len() - enter_index;
                    let hallway_cost = h-l+1;
                    new_state.cost += (enter_cost + hallway_cost) * tile.cost();
                    seen.insert(new_state);
                }
            }
        }

        for i in 0..self.rooms.len() {
            let room = &self.rooms[i];
            let room_hallway_index = (i + 1) * 2;
            let index_option = room.iter().rev().position(|t| t != &Tile::Empty);
            if index_option.is_none() {
                // Room is empty
                continue;
            }
            let index = room.len() - index_option.unwrap() - 1;
            let tile = room[index];

            let all_set_below = room[..index].iter().all(|t| t == &tile);
            let is_destination = i == tile.room_index();
            if all_set_below && is_destination {
                continue;
            }

            let mut valid_moves = Vec::new();
            // Check all to the right
            for j in room_hallway_index + 1..self.hallway.len() {
                if self.hallway[j] != Tile::Empty {
                    break;
                }
                // Can't stop outside rooms
                if j == 2 || j == 4 || j == 6 || j == 8 {
                    continue;
                }
                valid_moves.push(j);
            }
            // Check all to the left
            for j in (0..room_hallway_index).rev() {
                if self.hallway[j] != Tile::Empty {
                    break;
                }
                // Can't stop outside rooms
                if j == 2 || j == 4 || j == 6 || j == 8 {
                    continue;
                }
                valid_moves.push(j);
            }

            for m in valid_moves {
                let mut new_state = self.clone();
                new_state.hallway[m] = tile;
                new_state.rooms[i][index] = Tile::Empty;
                let out_of_room_cost = room.len() - index;
                let l = m.min(room_hallway_index);
                let h = m.max(room_hallway_index);
                let hallway_cost = (h - l).max(1);
                new_state.cost += (out_of_room_cost + hallway_cost) * tile.cost();
                seen.insert(new_state);
            }
        }
    }

    fn is_finished(self: &Self) -> bool {
        let a = self.rooms[0].iter().all(|t| t == &Tile::A);
        let b = self.rooms[1].iter().all(|t| t == &Tile::B);
        let c = self.rooms[2].iter().all(|t| t == &Tile::C);
        let d = self.rooms[3].iter().all(|t| t == &Tile::D);
        return a && b && c && d;
    }
}

fn part1(input: &str) -> usize {
    let initial_state = State::parse(input);
    let mut seen: HashSet<State> = HashSet::new();
    seen.insert(initial_state.clone());

    let mut last_size = 0;
    let mut curr_size = 1;

    while last_size != curr_size {
        for state in seen.clone() {
            state.step(&mut seen);
        }
        last_size = curr_size;
        curr_size = seen.len();
    }

    let mut min = usize::MAX;
    for state in seen {
        if state.is_finished() {
            min = min.min(state.cost);
        }
    }

    return min;
}

fn part2(input: &str) -> usize {
    let new_lines = "  #D#C#B#A#\n  #D#B#A#C#";
    let mut lines = input.trim().lines().collect::<Vec<&str>>();
    lines.insert(3, new_lines);
    return part1(&lines.join("\n"));
}

pub fn run() {
    let input = include_str!("../input/day23");
    println!("23.1: {:?}", part1(input));
    println!("23.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {

        assert_eq!(
            part1(
                "#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########"
            ),
            12521
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########"
            ),
            44169
        );
    }
}
