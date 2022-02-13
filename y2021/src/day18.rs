
#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Node {
    value: u32,
    depth: u32,
}

fn parse_number(input: &str) -> Vec<Node> {
    let mut nodes = Vec::new();
    let mut depth = 0;
    let mut buf = String::new();
    for c in input.trim().chars() {
        match c {
            '[' | ']' | ',' => {
                if !buf.is_empty() {
                    nodes.push(Node {
                        value: buf.parse::<u32>().unwrap(),
                        depth,
                    });
                    buf.clear();
                }

                if c == '[' {
                    depth += 1;
                } else if c == ']' {
                    depth -= 1;
                }
            },
            digit => {
                buf.push(digit);
            }
        }
    }
    return nodes;
}

fn add_snailfish_numbers(a: &Vec<Node>, b: &Vec<Node>) -> Vec<Node> {
    let mut res = Vec::new();
    for node in a {
        res.push(Node {
            value: node.value,
            depth: node.depth + 1,
        });
    }

    for node in b {
        res.push(Node {
            value: node.value,
            depth: node.depth + 1,
        });
    }

    return res;
}

fn reduce_snailfish_number(n: &mut Vec<Node>) -> Vec<Node> {
    for i in 0..n.len() {
        let curr = n[i];
        if curr.depth > 4 {
            let l = curr;
            let r = *n.get(i+1).expect("r bound break");

            n.remove(i);
            n.remove(i);
            n.insert(i, Node {
                value: 0,
                depth: curr.depth-1,
            });

            let len = n.len();

            if i > 0 {
                n[i-1].value += l.value;
            }
            if i < len - 1 {
                n[i+1].value += r.value;
            }

            return reduce_snailfish_number(n);
        }
    }
    for i in 0..n.len() {
        let curr = n[i];
        if curr.value >= 10 {
            let v = curr.value;
            let l = v / 2;
            let r = if v % 2 != 0 {
                l + 1
            } else {
                l
            };
            n.remove(i);
            n.insert(i, Node {
                value: l,
                depth: curr.depth + 1,
            });
            n.insert(i+1, Node {
                value: r,
                depth: curr.depth + 1,
            });

            return reduce_snailfish_number(n);
        }
    }

    return n.to_vec();
}

fn magnitude(n: &mut Vec<Node>) -> usize {
    let mut i = 0;
    let mut l = n.len();

    let mut c = 0;
    while l != 1 {
        c += 1;
        if c > 8 {
            break;
        }
        while i < l-1 {
            if n[i].depth == n[i+1].depth {
                let x = (3*n[i].value) + (2*n[i+1].value);
                n[i].value = x;
                n[i].depth -= 1;
                n.remove(i+1);
            } else {
                i += 1;
            }
            l = n.len();
        }
        i = 0;
    }

    return n[0].value as usize;
}

fn part1(input: &str) -> usize {
    let mut numbers = Vec::new();
    for line in input.trim().lines() {
        let res = parse_number(line);
        numbers.push(res);
    }

    let mut result = numbers[0].clone();
    for i in 1..numbers.len() {
        result = add_snailfish_numbers(&result, &numbers[i]);
        result = reduce_snailfish_number(&mut result);
    }

    return magnitude(&mut result);
}

fn part2(input: &str) -> usize {
    let mut numbers = Vec::new();
    for line in input.trim().lines() {
        let res = parse_number(line);
        numbers.push(res);
    }

    let mut max = usize::MIN;

    for i in 0..numbers.len() {
        for j in 0..numbers.len() {
            if i == j {
                continue;
            }
            let x = magnitude(&mut reduce_snailfish_number(&mut add_snailfish_numbers(&numbers[i], &numbers[j])));
            if x > max {
                max = x;
            }
        }
    }

    return max;
}

pub fn run() {
    let input = include_str!("../input/day18");
    println!("18.1: {:?}", part1(input));
    println!("18.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"), 4140);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"), 3993);
    }
}

