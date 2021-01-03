use std::fs;
use std::thread;
use std::sync::mpsc;
extern crate crypto;
use crypto::md5::Md5;
use crypto::digest::Digest;

fn finder(s: &str, n: usize, start: u64, end: u64) -> Option<u64> {
    let s = s.trim();
    let mut i: u64 = start;
    let mut md5 = Md5::new();

    loop {
        md5.input_str(s);
        md5.input_str(&i.to_string());
        let dig = md5.result_str();
        md5.reset();

        if dig.chars().take(n).all(|c| c == '0') {
            return Some(i);
        }
        if i >= end { break; }
        i += 1;
    }

    return None;
}

fn find_ith_suffix(s: &str, n: usize) -> u64 {
    let chunk = 1<<14;
    let procs = 4;

    let mut res: u64 = std::u64::MAX;
    let mut done: bool = false;

    let mut chunk_iter = (0..).map(|x| x*chunk).take_while(|x| x < &std::u64::MAX);

    loop {
        let mut children = vec![];
        for _ in 0..procs {
            let s = s.to_owned();
            let b = chunk_iter.next().unwrap();
            children.push(thread::spawn(move ||
                    finder(&s, n, b, b+chunk)));
        }

        for child in children.into_iter() {
            if let Some(x) = &child.join().unwrap() {
                res = res.min(*x);
                done = true;
            }
        }

        if done { break; }
    }

    return res;
}

pub fn run () {
    let line = &fs::read_to_string("./input/day04.txt")
        .expect("Could not open input")[..];

    println!("4a: {}", find_ith_suffix(line, 5));
    println!("4b: {}", find_ith_suffix(line, 6));
}
