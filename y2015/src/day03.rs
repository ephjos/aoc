use std::fs;
use std::collections::HashMap;

fn a (s: &str) {
    let mut hm = HashMap::new();
    let mut curr = Vec2D{ x: 0, y: 0 };

    hm.insert(curr, 1);

    for c in s.chars() {
        match c {
            '^' => curr.y += 1,
            'v' => curr.y -= 1,
            '>' => curr.x += 1,
            '<' => curr.x -= 1,
            _ => continue,
        }

        let count = hm.entry(curr).or_insert(0);
        *count += 1;
    }

    println!("3a: {:#?}", hm.keys().len());
}


fn b (s: &str) {
    let mut hm = HashMap::new();
    let mut santa = Vec2D{ x: 0, y: 0 };
    let mut robot = Vec2D{ x: 0, y: 0 };

    hm.insert(santa, 1);

    for (i,c) in s.chars().enumerate() {
        if i % 2 == 0 {
            santa.move_char(c);
            let count = hm.entry(santa).or_insert(0);
            *count += 1;
        } else {
            robot.move_char(c);
            let count = hm.entry(robot).or_insert(0);
            *count += 1;
        }
    }

    println!("3b: {:#?}", hm.keys().len());
}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
struct Vec2D {
    x: i64,
    y: i64,
}

impl Vec2D {
    fn move_char(&mut self, c: char) {
        match c {
            '^' => self.y += 1,
            'v' => self.y -= 1,
            '>' => self.x += 1,
            '<' => self.x -= 1,
            _ => (),
        }
    }
}

pub fn run () {
    let cmds = fs::read_to_string("./input/day03.txt")
        .expect("Could not open input");

    a(&cmds[..]);
    b(&cmds[..]);
}
