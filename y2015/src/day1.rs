use std::fs;

fn a(s: &String) {
    let mut x = 0;
    for c in s.chars() {
        if c == '(' {
            x += 1;
        } else if c == ')' {
            x -= 1;
        }
    }

    println!("1a: {}", x);
}

fn b(s: &String) {
    let mut x: i32 = 0;
    let mut res: usize = 0;
    for (i,c) in s.chars().enumerate() {
        if c == '(' {
            x += 1;
        } else if c == ')' {
            x -= 1;
        }

        if x < 0 {
            res = i;
            break;
        }
    }

    println!("1b: {}", res+1);
}

pub fn run () {
    let s = fs::read_to_string("./input/day1.txt")
        .expect("Could not read input");

    a(&s);
    b(&s);
}
