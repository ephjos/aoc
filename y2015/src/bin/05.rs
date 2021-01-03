use std::fs;

y2015::main! {
    let strings = fs::read_to_string("./input/day05.txt").unwrap();
    let a_nice = strings.lines().filter(|s| {
        let vowels = ['a','e','i','o','u'];
        let illegal = [
            String::from("ab"),
            String::from("cd"),
            String::from("pq"),
            String::from("xy"),
        ];
        let mut is_dub = false;
        let mut vowel_count = 0;
        let mut prev: char = '\0';

        for c in s.chars() {
            if vowels.contains(&c) {
                vowel_count += 1;
            }

            if illegal.contains(&format!("{}{}",prev,c)) {
                return false;
            }

            is_dub |= prev == c;

            prev = c;
        }

        if vowel_count < 3 { return false; }
        if !is_dub { return false; }

        return true;
    }).collect::<Vec<_>>();

    let b_nice = strings.lines().filter(|s| {
        let cs: Vec<char> = s.chars().collect::<Vec<char>>();

        let mut any_pair = false;
        let mut split_pair = false;

        for i in 0..cs.len()-2 {
            let ts = &format!("{}{}",cs[i],cs[i+1]);
            any_pair |= s[i+2..].contains(ts);

            split_pair |= cs[i] == cs[i+2];
        }

        return any_pair && split_pair;
    }).collect::<Vec<_>>();

    (a_nice.len(), b_nice.len())
}
