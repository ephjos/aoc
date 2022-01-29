
#[derive(Debug, Eq, PartialEq, Clone)]
struct Packet {
    version: u64,
    type_id: u64,
    literal: Option<u64>,
    length_type_id: Option<u64>,
    packets: Option<Vec<Packet>>,
}

impl Packet {
    fn sum(self: &Self) -> u64 {
        let mut res = 0;

        if let Some(children) = &self.packets {
            for c in children {
                res += c.calculate_value();
            }
        }

        return res;
    }

    fn product(self: &Self) -> u64 {
        let mut res = 1;

        if let Some(children) = &self.packets {
            for c in children {
                res *= c.calculate_value();
            }
        }

        return res;
    }

    fn minimum(self: &Self) -> u64 {
        let mut res = u64::MAX;

        if let Some(children) = &self.packets {
            for c in children {
                let v = c.calculate_value();
                if v < res {
                    res = v;
                }
            }
        }

        return res;
    }

    fn maximum(self: &Self) -> u64 {
        let mut res = u64::MIN;

        if let Some(children) = &self.packets {
            for c in children {
                let v = c.calculate_value();
                if v > res {
                    res = v;
                }
            }
        }

        return res;
    }

    fn greater(self: &Self) -> u64 {
        let mut res = 0;

        if let Some(children) = &self.packets {
            let l = &children[0];
            let r = &children[1];

            if l.calculate_value() > r.calculate_value() {
                res = 1;
            }
        }

        return res;
    }

    fn lesser(self: &Self) -> u64 {
        let mut res = 0;

        if let Some(children) = &self.packets {
            let l = &children[0];
            let r = &children[1];

            if l.calculate_value() < r.calculate_value() {
                res = 1;
            }
        }

        return res;
    }

    fn equal(self: &Self) -> u64 {
        let mut res = 0;

        if let Some(children) = &self.packets {
            let l = &children[0];
            let r = &children[1];

            if l.calculate_value() == r.calculate_value() {
                res = 1;
            }
        }

        return res;
    }

    fn calculate_value(self: &Self) -> u64 {
        return match self.type_id {
            0 => self.sum(),
            1 => self.product(),
            2 => self.minimum(),
            3 => self.maximum(),
            4 => self.literal.unwrap(),
            5 => self.greater(),
            6 => self.lesser(),
            7 => self.equal(),
            _ => 0,
        };
    }
}

fn hex_string_to_bit_vec(input: &str) -> String {
    let mut res = String::new();
    input.trim().chars().for_each(|c| {
        res.push_str(match c {
            '0' => "0000",
            '1' => "0001",
            '2' => "0010",
            '3' => "0011",
            '4' => "0100",
            '5' => "0101",
            '6' => "0110",
            '7' => "0111",
            '8' => "1000",
            '9' => "1001",
            'A' => "1010",
            'B' => "1011",
            'C' => "1100",
            'D' => "1101",
            'E' => "1110",
            'F' => "1111",
            _ => "",
        });
    });
    return res;
}

fn bits_to_u64(input: &str) -> u64 {
    return u64::from_str_radix(input, 2).unwrap();
}

fn parse_literal(input: &str) -> (&str, u64) {
    let mut res = String::new();
    let mut i = 0;
    while i < input.len() {
        res.push_str(&input[i+1..i+5]);
        if input.chars().nth(i).unwrap() == '0' {
            i += 5;
            break;
        }
        i += 5;
    }
    return (&input[i..], bits_to_u64(&res));
}

fn parse_packet(input: &str) -> (&str, Packet) {
    let version = bits_to_u64(&input[..3]);
    let type_id = bits_to_u64(&input[3..6]);

    if type_id == 4 {
        // Literal
        let (input_left, literal) = parse_literal(&input[6..]);
        return (input_left, Packet {
            version,
            type_id,
            literal: Some(literal),
            length_type_id: None,
            packets: None,
        })
    }

    // Operator
    let length_type_id: u64 = if input.chars().nth(6).unwrap() == '0' {
        0
    } else {
        1
    };
    let mut packets = Vec::new();
    let mut input_left;

    if length_type_id == 0 {
        let mut bits_left = bits_to_u64(&input[7..22]);

        input_left = &input[22..];
        while bits_left > 0 {
            let res = parse_packet(input_left);
            let next_input_left = res.0;

            bits_left -= input_left.len() as u64 - next_input_left.len() as u64;
            input_left = next_input_left;

            let packet = res.1;
            packets.push(packet);
        }
    } else {
        let packets_left = bits_to_u64(&input[7..18]);

        input_left = &input[18..];
        for _ in 0..packets_left {
            let res = parse_packet(input_left);
            input_left = res.0;
            packets.push(res.1);
        }
    }

    return (input_left, Packet {
        version,
        type_id,
        length_type_id: Some(length_type_id),
        literal: None,
        packets: Some(packets),
    });
}

fn sum_version_numbers(packet: &Packet) -> u64 {
    let mut res = packet.version;
    if let Some(children) = &packet.packets {
        for c in children {
            res += sum_version_numbers(&c);
        }
    }
    return res;
}

fn part1(input: &str) -> u64 {
    let bit_vec = &hex_string_to_bit_vec(input);
    let (_, packet) = parse_packet(&bit_vec);

    return sum_version_numbers(&packet);
}

fn part2(input: &str) -> u64 {
    let bit_vec = &hex_string_to_bit_vec(input);
    let (_, packet) = parse_packet(&bit_vec);

    return packet.calculate_value();
}

pub fn run() {
    let input = include_str!("../input/day16");
    println!("16.1: {:?}", part1(input));
    println!("16.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("D2FE28"), 6);
        assert_eq!(part1("38006F45291200"), 9);
        assert_eq!(part1("EE00D40C823060"), 14);
        assert_eq!(part1("8A004A801A8002F478"), 16);
        assert_eq!(part1("620080001611562C8802118E34"), 12);
        assert_eq!(part1("C0015000016115A2E0802F182340"), 23);
        assert_eq!(part1("A0016C880162017C3686B18A3D4780"), 31);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("C200B40A82"), 3);
        assert_eq!(part2("04005AC33890"), 54);
        assert_eq!(part2("880086C3E88112"), 7);
        assert_eq!(part2("CE00C43D881120"), 9);
        assert_eq!(part2("D8005AC2A8F0"), 1);
        assert_eq!(part2("F600BC2D8F"), 0);
        assert_eq!(part2("9C005AC2F8F0"), 0);
        assert_eq!(part2("9C0141080250320F1802104A08"), 1);
    }
}

