use std::fs;
use md5;

pub fn run () {
    let line = fs::read_to_string("./input/day04.txt")
        .expect("Could not open input");

    let s = line.trim();
    let mut a = false;
    let mut b = false;
    let mut i: u64 = 0;
    loop {
        let dig = md5::compute([s, &i.to_string()[..]].join(""));
        let dig_s = format!("{:x}",dig);
        if !a && dig_s.starts_with("00000") {
            println!("4a: {}", i);
            a = true;
            if b { break }
        }
        if !b && dig_s.starts_with("000000") {
            println!("4b: {}", i);
            b = true;
            if a { break }
        }
        i += 1;
    }
}
