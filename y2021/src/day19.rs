use std::collections::{HashSet, HashMap};

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Clone, Copy)]
struct Point3D {
    x: isize,
    y: isize,
    z: isize,
}



impl std::ops::Add<Point3D> for Point3D {
    type Output = Point3D;

    fn add(self, rhs: Point3D) -> Point3D {
        Point3D {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
            z: self.z + rhs.z,
        }
    }
}
impl std::ops::Sub<Point3D> for Point3D {
    type Output = Point3D;

    fn sub(self, rhs: Point3D) -> Point3D {
        Point3D {
            x: self.x - rhs.x,
            y: self.y - rhs.y,
            z: self.z - rhs.z,
        }
    }
}

#[derive(Debug, Clone)]
struct Scanner3D {
    rotations: Vec<HashSet<Point3D>>,
    matched_rotation_index: usize,
    matched_position: Point3D,
}

impl Scanner3D {
    pub fn parse(input: &str) -> Self {
        let beacons = input
            .trim()
            .lines()
            .skip(1)
            .map(|l| {
                let nums = l
                    .splitn(3, ",")
                    .map(|t| t.parse::<isize>().unwrap())
                    .collect::<Vec<isize>>();
                assert_eq!(nums.len(), 3);
                Point3D {
                    x: nums[0],
                    y: nums[1],
                    z: nums[2],
                }
            })
            .collect::<Vec<Point3D>>();

        let mut rotations: Vec<HashSet<Point3D>> = Vec::new();
        //+x+y+z
        rotations.push(HashSet::from_iter(beacons.to_vec()));

        //-x-y+z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: -p.y,
                    z: p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));

        //-x+y-z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: p.y,
                    z: -p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+x-y-z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: -p.y,
                    z: -p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-x+z+y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: p.z,
                    z: p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+x-z+y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: -p.z,
                    z: p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+x+z-y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: p.z,
                    z: -p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-x-z-y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: -p.z,
                    z: -p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-y+x+z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: p.x,
                    z: p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+y-x+z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: -p.x,
                    z: p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+y+x-z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: p.x,
                    z: -p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-y-x-z
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: -p.x,
                    z: -p.z,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+y+z+x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: p.z,
                    z: p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-y-z+x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: -p.z,
                    z: p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-y+z-x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: p.z,
                    z: -p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+y-z-x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: -p.z,
                    z: -p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+z+x+y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: p.x,
                    z: p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-z-x+y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: -p.x,
                    z: p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-z+x-y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: p.x,
                    z: -p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+z-x-y
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: -p.x,
                    z: -p.y,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-z+y+x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: p.y,
                    z: p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+z-y+x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: -p.y,
                    z: p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //+z+y-x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: p.y,
                    z: -p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));
        //-z-y-x
        rotations.push(HashSet::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: -p.y,
                    z: -p.x,
                })
                .collect::<HashSet<Point3D>>(),
        ));

        Scanner3D {
            rotations,
            matched_rotation_index: 0,
            matched_position: Point3D { x: 0, y: 0, z: 0 },
        }
    }

    pub fn overlap(self: &Self, other: &mut Self) -> bool {
        let base_points = &self.rotations[self.matched_rotation_index];
        for r in 0..24 {
            let target_points = &other.rotations[r];
            let mut distances = Vec::new();
            for base in base_points {
                for target in target_points {
                    distances.push((*base+self.matched_position)-*target);
                }
            }
            let frequencies = distances
                .iter()
                .copied()
                .fold(HashMap::new(), |mut map, val|{
                    map.entry(val)
                        .and_modify(|frq|*frq+=1)
                        .or_insert(1);
                    map
                });
            let mut max_p = Point3D{x:0,y:0,z:0};
            let mut max_f = 0;
            for (p, f) in frequencies {
                if f > max_f {
                    max_p = p;
                    max_f = f;
                }
            }
            if max_f >= 12 {
                other.matched_rotation_index = r;
                other.matched_position = max_p;
                return true;
            }
        }
        return false;
    }
}

fn part1(input: &str) -> usize {
    let mut scanners = input
        .trim()
        .split("\n\n")
        .map(|block| Scanner3D::parse(block))
        .collect::<Vec<Scanner3D>>();
    let mut aligned = vec![scanners.remove(0)];

    while !scanners.is_empty() {
        for i in 0..aligned.len() {
            if scanners.is_empty() {
                break;
            }
            let s = aligned[i].clone();
            let mut overlap = false;
            for j in 0..scanners.len() {
                overlap = s.overlap(&mut scanners[j]);
                if overlap {
                    aligned.push(scanners.remove(j));
                    break;
                }
            }
            if overlap {
                break;
            }
        }
    }


    let mut all_beacons = HashSet::new();

    for scanner in aligned {
        let pos = scanner.matched_position;
        for beacon in &scanner.rotations[scanner.matched_rotation_index] {
            all_beacons.insert(*beacon + pos);
        }
    }

    return all_beacons.len();
}

fn part2(input: &str) -> isize {
    let mut scanners = input
        .trim()
        .split("\n\n")
        .map(|block| Scanner3D::parse(block))
        .collect::<Vec<Scanner3D>>();
    let mut aligned = vec![scanners.remove(0)];

    while !scanners.is_empty() {
        for i in 0..aligned.len() {
            if scanners.is_empty() {
                break;
            }
            let s = aligned[i].clone();
            let mut overlap = false;
            for j in 0..scanners.len() {
                overlap = s.overlap(&mut scanners[j]);
                if overlap {
                    aligned.push(scanners.remove(j));
                    break;
                }
            }
            if overlap {
                break;
            }
        }
    }

    let mut m_distances = Vec::new();
    for i in 0..aligned.len() {
        for j in 0..aligned.len() {
            if i == j {
                continue;
            }
            let a = &aligned[i];
            let b = &aligned[j];
            let diff = a.matched_position - b.matched_position;
            let dist = diff.x.abs() + diff.y.abs() + diff.z.abs();

            m_distances.push(dist);
        }
    }

    return *m_distances.iter().max().unwrap();
}

pub fn run() {
        assert_eq!(
            part1(
                "--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14"
            ),
            79
        );
    let input = include_str!("../input/day19");
    println!("19.1: {:?}", part1(input));
    println!("19.2: {:?}", part2(input));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(
                "--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14"
            ),
            79
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(
                "--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14"
            ),
            3621
        );
    }
}
