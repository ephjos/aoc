use std::fs;

fn a(s: &String) -> i32 {
    let mut x = 0;
    for c in s.chars() {
        if c == '(' {
            x += 1;
        } else if c == ')' {
            x -= 1;
        }
    }

    return x;
}

fn b(s: &String) -> i32 {
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

    return (res+1) as i32;
}

y2015::main! {
    let s = &fs::read_to_string("./input/day01.txt")
        .expect("Could not read input");

    (a(s), b(s))
}
