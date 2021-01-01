use std::fs;

fn a () {}

fn b () {}

pub fn run () {
    let content = fs::read_to_string("./input/dayNN.txt")
        .expect("Could not open input");
}
