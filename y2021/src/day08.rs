use super::com::*;

fn part1(input: &str) -> isize {
    fn word<'a>() -> impl Parser<'a, String> {
        whitespace_wrap(one_or_more(pred(any_char, |c| c.is_alphanumeric())))
            .map(|cs| cs.into_iter().collect())
    }

    fn words<'a>() -> impl Parser<'a, Vec<String>> {
        one_or_more(word())
    }

    fn entry_list_parser<'a>() -> impl Parser<'a, (Vec<String>, Vec<String>)> {
        pair(left(words(), whitespace_wrap(match_literal("|"))), words())
    }

    let entries: Vec<(Vec<String>, Vec<String>)> = input
        .lines()
        .map(|l| entry_list_parser().parse(l).unwrap().1)
        .collect();

    let mut count = 0;
    for (_unique, output) in entries {
        for o in output {
            let l = o.len();
            if l == 2 || l == 3 || l == 4 || l == 7 {
                count += 1;
            }
        }
    }

    return count;
}

fn part2(input: &str) -> usize {
    fn word<'a>() -> impl Parser<'a, String> {
        whitespace_wrap(one_or_more(pred(any_char, |c| c.is_alphanumeric())))
            .map(|cs| cs.into_iter().collect())
    }

    fn words<'a>() -> impl Parser<'a, Vec<String>> {
        one_or_more(word())
    }

    fn entry_list_parser<'a>() -> impl Parser<'a, (Vec<String>, Vec<String>)> {
        pair(left(words(), whitespace_wrap(match_literal("|"))), words())
    }

    let mut entries: Vec<(Vec<String>, Vec<String>)> = input
        .lines()
        .map(|l| entry_list_parser().parse(l).unwrap().1)
        .collect();

    let mut result = 0;
    for (u, outputs) in entries {
        let u2 = u.iter().find(|w| w.len() == 2).unwrap();
        let u3 = u.iter().find(|w| w.len() == 3).unwrap();
        let u4 = u.iter().find(|w| w.len() == 4).unwrap();
        let u5s: Vec<String> = u.iter().filter(|w| w.len() == 5).cloned().collect();
        let u6s: Vec<String> = u.iter().filter(|w| w.len() == 6).cloned().collect();
        let u7 = u.iter().find(|w| w.len() == 7).unwrap();

        let top = u3.chars().find(|c| !u2.contains(*c)).unwrap();
        let middle = u4.chars().find(|c| u5s[0].contains(*c) && u5s[1].contains(*c) && u5s[2].contains(*c)).unwrap();
        let top_left = u4.chars().find(|c| *c != middle && !u2.contains(*c)).unwrap();
        let top_right = u2.chars().find(|c| !(u6s[0].contains(*c) && u6s[1].contains(*c) && u6s[2].contains(*c))).unwrap();
        let bottom_right = u2.chars().find(|c| *c != top_right).unwrap();
        let the_9 = u6s.iter().find(|s| s.contains(top_right) && s.contains(bottom_right)).unwrap();
        let bottom = the_9.chars().find(|c| *c != top && *c != middle && *c != top_left && *c != top_right && *c != bottom_right).unwrap();
        let bottom_left = u7.chars().find(|c| *c != top && *c != middle && *c != top_left && *c != top_right && *c != bottom_right && *c != bottom).unwrap();
        /*
        println!(" {}{}{}{} ", top,top,top,top);
        println!("{}    {}", top_left,top_right);
        println!("{}    {}", top_left,top_right);
        println!(" {}{}{}{} ", middle,middle,middle,middle);
        println!("{}    {}", bottom_left,bottom_right);
        println!("{}    {}", bottom_left,bottom_right);
        println!(" {}{}{}{} ", bottom,bottom,bottom,bottom);
        println!("");
        */

        let mut digits = String::new();
        for o in outputs {
            let l = o.len();
            match l {
                2 => digits.push('1'),
                3 => digits.push('7'),
                4 => digits.push('4'),
                5 => {
                    if o.contains(top_right) && o.contains(bottom_right) {
                        digits.push('3');
                    } else if o.contains(top_right) && o.contains(bottom_left) {
                        digits.push('2');
                    } else {
                        digits.push('5');
                    }
                },
                6 => {
                    if !o.contains(middle) {
                        digits.push('0');
                    } else if o.contains(top_right) {
                        digits.push('9');
                    } else {
                        digits.push('6');
                    }
                },
                7 => digits.push('8'),
                _ => unreachable!(),
            }
        }
        let output = digits.parse::<usize>().unwrap();
        result += output;
    }

    return result;
}

pub fn run() {
    let input = include_str!("../input/day08");
    println!("08.1: {:?}", part1(input));
    println!("08.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"
        ),
            26
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |cdfeb fcadb cdfeb cdbaf"),
            5353
        );
        assert_eq!(
            part2(
                "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"
        ),
            61229
        );
    }
}
