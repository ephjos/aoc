#[derive(Clone, Copy)]
struct Range3D {
    on: isize,
    x0: isize,
    y0: isize,
    z0: isize,
    x1: isize,
    y1: isize,
    z1: isize,
}

impl std::fmt::Debug for Range3D {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "({} ({},{},{}) ({},{},{}) [{}])\n",
            self.on,
            self.x0,
            self.y0,
            self.z0,
            self.x1,
            self.y1,
            self.z1,
            self.volume()
        )
    }
}

impl Range3D {
    fn parse(line: &str) -> Range3D {
        let tr = line.trim();
        let on = &tr[0..3] == "on ";
        let num_parts = &tr[3..].trim().splitn(3, ",").collect::<Vec<&str>>();
        let vals = num_parts
            .iter()
            .map(|p| {
                return p[2..]
                    .splitn(2, "..")
                    .map(|n| n.parse::<isize>().unwrap())
                    .collect::<Vec<isize>>();
            })
            .collect::<Vec<Vec<isize>>>();

        let x0 = vals[0][0];
        let y0 = vals[1][0];
        let z0 = vals[2][0];
        let x1 = vals[0][1];
        let y1 = vals[1][1];
        let z1 = vals[2][1];

        Range3D {
            on: on as isize,
            x0,
            y0,
            z0,
            x1,
            y1,
            z1,
        }
    }

    fn intersect(self: &Self, other: &Range3D) -> Option<Range3D> {
        if (self.x0 <= other.x1 && self.x1 >= other.x0)
            && (self.y0 <= other.y1 && self.y1 >= other.y0)
            && (self.z0 <= other.z1 && self.z1 >= other.z0)
        {
            let on = other.on;

            let x0 = self.x0.max(other.x0);
            let y0 = self.y0.max(other.y0);
            let z0 = self.z0.max(other.z0);

            let x1 = self.x1.min(other.x1);
            let y1 = self.y1.min(other.y1);
            let z1 = self.z1.min(other.z1);

            return Some(Range3D {
                on,
                x0,
                y0,
                z0,
                x1,
                y1,
                z1,
            });
        }

        None
    }

    fn volume(self: &Self) -> isize {
        return self.on
            * (self.x1 - self.x0 + 1)
            * (self.y1 - self.y0 + 1)
            * (self.z1 - self.z0 + 1);
    }
}

fn part1(input: &str) -> isize {
    let ranges = input
        .trim()
        .lines()
        .map(|l| Range3D::parse(l))
        .collect::<Vec<Range3D>>();
    let mut out: Vec<Range3D> = Vec::new();

    for i in 0..ranges.len() {
        let mut curr = &mut ranges[i].clone();
        if curr.x0 > 50 || curr.x1 < -50 ||
        curr.y0 > 50 || curr.y1 < -50 ||
        curr.z0 > 50 || curr.z1 < -50 {
            continue;
        }
        curr.x0 = curr.x0.max(-50);
        curr.y0 = curr.y0.max(-50);
        curr.z0 = curr.z0.max(-50);

        curr.x1 = curr.x1.min(50);
        curr.y1 = curr.y1.min(50);
        curr.z1 = curr.z1.min(50);

        let mut buf = Vec::new();
        for j in 0..out.len() {
            let other = &out[j];
            if !((curr.x0 <= other.x1 && curr.x1 >= other.x0)
                && (curr.y0 <= other.y1 && curr.y1 >= other.y0)
                && (curr.z0 <= other.z1 && curr.z1 >= other.z0))
            {
                continue;
            }

            let x0 = curr.x0.max(other.x0);
            let y0 = curr.y0.max(other.y0);
            let z0 = curr.z0.max(other.z0);

            let x1 = curr.x1.min(other.x1);
            let y1 = curr.y1.min(other.y1);
            let z1 = curr.z1.min(other.z1);
            buf.push(Range3D {
                on: -1 * other.on,
                x0,
                y0,
                z0,
                x1,
                y1,
                z1,
            });
        }
        if curr.on == 1 {
            out.push(*curr);
        }
        out.append(&mut buf);
    }

    return out.iter().map(|o| o.volume()).sum();
}

fn part2(input: &str) -> isize {
    let ranges = input
        .trim()
        .lines()
        .map(|l| Range3D::parse(l))
        .collect::<Vec<Range3D>>();
    let mut out: Vec<Range3D> = Vec::new();

    for i in 0..ranges.len() {
        let mut curr = &mut ranges[i].clone();
        let mut buf = Vec::new();
        for j in 0..out.len() {
            let other = &out[j];
            if !((curr.x0 <= other.x1 && curr.x1 >= other.x0)
                && (curr.y0 <= other.y1 && curr.y1 >= other.y0)
                && (curr.z0 <= other.z1 && curr.z1 >= other.z0))
            {
                continue;
            }

            let x0 = curr.x0.max(other.x0);
            let y0 = curr.y0.max(other.y0);
            let z0 = curr.z0.max(other.z0);

            let x1 = curr.x1.min(other.x1);
            let y1 = curr.y1.min(other.y1);
            let z1 = curr.z1.min(other.z1);
            buf.push(Range3D {
                on: -1 * other.on,
                x0,
                y0,
                z0,
                x1,
                y1,
                z1,
            });
        }
        if curr.on == 1 {
            out.push(*curr);
        }
        out.append(&mut buf);
    }

    return out.iter().map(|o| o.volume()).sum();
}

pub fn run() {
    let input = include_str!("../input/day22");
    println!("22.1: {:?}", part1(input));
    println!("22.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10"
            ),
            39
        );
        assert_eq!(
            part1(
                "on x=-20..26,y=-36..17,z=-47..7
on x=-20..33,y=-21..23,z=-26..28
on x=-22..28,y=-29..23,z=-38..16
on x=-46..7,y=-6..46,z=-50..-1
on x=-49..1,y=-3..46,z=-24..28
on x=2..47,y=-22..22,z=-23..27
on x=-27..23,y=-28..26,z=-21..29
on x=-39..5,y=-6..47,z=-3..44
on x=-30..21,y=-8..43,z=-13..34
on x=-22..26,y=-27..20,z=-29..19
off x=-48..-32,y=26..41,z=-47..-37
on x=-12..35,y=6..50,z=-50..-2
off x=-48..-32,y=-32..-16,z=-15..-5
on x=-18..26,y=-33..15,z=-7..46
off x=-40..-22,y=-38..-28,z=23..41
on x=-16..35,y=-41..10,z=-47..6
off x=-32..-23,y=11..30,z=-14..3
on x=-49..-5,y=-3..45,z=-29..18
off x=18..30,y=-20..-8,z=-3..13
on x=-41..9,y=-7..43,z=-33..15
on x=-54112..-39298,y=-85059..-49293,z=-27449..7877
on x=967..23432,y=45373..81175,z=27513..53682"
            ),
            590784
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "on x=-5..47,y=-31..22,z=-19..33
on x=-44..5,y=-27..21,z=-14..35
on x=-49..-1,y=-11..42,z=-10..38
on x=-20..34,y=-40..6,z=-44..1
off x=26..39,y=40..50,z=-2..11
on x=-41..5,y=-41..6,z=-36..8
off x=-43..-33,y=-45..-28,z=7..25
on x=-33..15,y=-32..19,z=-34..11
off x=35..47,y=-46..-34,z=-11..5
on x=-14..36,y=-6..44,z=-16..29
on x=-57795..-6158,y=29564..72030,z=20435..90618
on x=36731..105352,y=-21140..28532,z=16094..90401
on x=30999..107136,y=-53464..15513,z=8553..71215
on x=13528..83982,y=-99403..-27377,z=-24141..23996
on x=-72682..-12347,y=18159..111354,z=7391..80950
on x=-1060..80757,y=-65301..-20884,z=-103788..-16709
on x=-83015..-9461,y=-72160..-8347,z=-81239..-26856
on x=-52752..22273,y=-49450..9096,z=54442..119054
on x=-29982..40483,y=-108474..-28371,z=-24328..38471
on x=-4958..62750,y=40422..118853,z=-7672..65583
on x=55694..108686,y=-43367..46958,z=-26781..48729
on x=-98497..-18186,y=-63569..3412,z=1232..88485
on x=-726..56291,y=-62629..13224,z=18033..85226
on x=-110886..-34664,y=-81338..-8658,z=8914..63723
on x=-55829..24974,y=-16897..54165,z=-121762..-28058
on x=-65152..-11147,y=22489..91432,z=-58782..1780
on x=-120100..-32970,y=-46592..27473,z=-11695..61039
on x=-18631..37533,y=-124565..-50804,z=-35667..28308
on x=-57817..18248,y=49321..117703,z=5745..55881
on x=14781..98692,y=-1341..70827,z=15753..70151
on x=-34419..55919,y=-19626..40991,z=39015..114138
on x=-60785..11593,y=-56135..2999,z=-95368..-26915
on x=-32178..58085,y=17647..101866,z=-91405..-8878
on x=-53655..12091,y=50097..105568,z=-75335..-4862
on x=-111166..-40997,y=-71714..2688,z=5609..50954
on x=-16602..70118,y=-98693..-44401,z=5197..76897
on x=16383..101554,y=4615..83635,z=-44907..18747
off x=-95822..-15171,y=-19987..48940,z=10804..104439
on x=-89813..-14614,y=16069..88491,z=-3297..45228
on x=41075..99376,y=-20427..49978,z=-52012..13762
on x=-21330..50085,y=-17944..62733,z=-112280..-30197
on x=-16478..35915,y=36008..118594,z=-7885..47086
off x=-98156..-27851,y=-49952..43171,z=-99005..-8456
off x=2032..69770,y=-71013..4824,z=7471..94418
on x=43670..120875,y=-42068..12382,z=-24787..38892
off x=37514..111226,y=-45862..25743,z=-16714..54663
off x=25699..97951,y=-30668..59918,z=-15349..69697
off x=-44271..17935,y=-9516..60759,z=49131..112598
on x=-61695..-5813,y=40978..94975,z=8655..80240
off x=-101086..-9439,y=-7088..67543,z=33935..83858
off x=18020..114017,y=-48931..32606,z=21474..89843
off x=-77139..10506,y=-89994..-18797,z=-80..59318
off x=8476..79288,y=-75520..11602,z=-96624..-24783
on x=-47488..-1262,y=24338..100707,z=16292..72967
off x=-84341..13987,y=2429..92914,z=-90671..-1318
off x=-37810..49457,y=-71013..-7894,z=-105357..-13188
off x=-27365..46395,y=31009..98017,z=15428..76570
off x=-70369..-16548,y=22648..78696,z=-1892..86821
on x=-53470..21291,y=-120233..-33476,z=-44150..38147
off x=-93533..-4276,y=-16170..68771,z=-104985..-24507"
            ),
            2758514936282235
        );
    }
}
