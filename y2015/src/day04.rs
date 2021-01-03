use std::fs;
use std::thread;
use std::sync::mpsc;
extern crate crypto;
use crypto::md5::Md5;
use crypto::digest::Digest;

//
// Refinements from:
//   https://gist.github.com/coriolinus/f94748a0d232d32e4eb1
//
//   Using channels saves ~0.2s
//

enum Cmd {
    Chunk(u64, u64),
}

fn finder(
    s: &str, n: usize,
    res_tx: mpsc::Sender<(mpsc::Sender<Cmd>, Option<u64>)>) {

    let (cmd_tx, cmd_rx) = mpsc::channel();

    res_tx.send((cmd_tx.clone(), None)).unwrap();

    let mut start;
    let mut end;
    let s = s.trim();
    let mut md5 = Md5::new();

    loop {
        match cmd_rx.recv().unwrap() {
            Cmd::Chunk(a,b) => {
                start = a;
                end = b;
            },
        }

        for i in start..end {
            md5.input_str(s);
            md5.input_str(&i.to_string());
            let dig = md5.result_str();
            md5.reset();

            if dig.chars().take(n).all(|c| c == '0') {
                res_tx.send((cmd_tx.clone(), Some(i))).unwrap();
            }
        }

        match res_tx.send((cmd_tx.clone(), None)) {
            _ => (),
        }
    }
}

fn find_ith_suffix(s: &str, n: usize) -> u64 {
    let chunk = 1<<12;
    let threads = 8;

    let mut res: u64 = 0;
    let chunk_iter = (0..).map(|x| (x*chunk, (x+1)*chunk));
    let (res_tx, res_rx) = mpsc::channel();

    for _ in 0..threads {
        let s = s.to_owned();
        let res_tx = res_tx.clone();
        thread::spawn(move ||
                finder(&s, n, res_tx));
    }

    for chunk in chunk_iter {
        let (cmd_tx, val) = res_rx.recv().unwrap();

        if let Some(x) = val {
            res = x;
            for _ in 1..threads {
                let (_, val) = res_rx.recv().unwrap();
                if let Some(y) = val {
                    res = res.min(y);
                }
            }
            break;
        } else {
            cmd_tx.send(Cmd::Chunk(chunk.0, chunk.1)).unwrap();
        }
    }

    return res;
}

pub fn run () {
    let line = &fs::read_to_string("./input/day04.txt")
        .expect("Could not open input")[..];

    println!("4a: {}", find_ith_suffix(line, 5));
    println!("4b: {}", find_ith_suffix(line, 6));
}
