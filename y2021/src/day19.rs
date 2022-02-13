use std::collections::{HashMap, HashSet};

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
    rotations: Vec<Vec<Point3D>>,
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

        let mut rotations: Vec<Vec<Point3D>> = Vec::new();
        //+x+y+z
        rotations.push(Vec::from_iter(beacons.to_vec()));

        //-x-y+z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: -p.y,
                    z: p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));

        //-x+y-z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: p.y,
                    z: -p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+x-y-z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: -p.y,
                    z: -p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-x+z+y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: p.z,
                    z: p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+x-z+y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: -p.z,
                    z: p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+x+z-y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.x,
                    y: p.z,
                    z: -p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-x-z-y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.x,
                    y: -p.z,
                    z: -p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-y+x+z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: p.x,
                    z: p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+y-x+z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: -p.x,
                    z: p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+y+x-z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: p.x,
                    z: -p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-y-x-z
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: -p.x,
                    z: -p.z,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+y+z+x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: p.z,
                    z: p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-y-z+x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: -p.z,
                    z: p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-y+z-x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.y,
                    y: p.z,
                    z: -p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+y-z-x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.y,
                    y: -p.z,
                    z: -p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+z+x+y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: p.x,
                    z: p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-z-x+y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: -p.x,
                    z: p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-z+x-y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: p.x,
                    z: -p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+z-x-y
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: -p.x,
                    z: -p.y,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-z+y+x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: p.y,
                    z: p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+z-y+x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: -p.y,
                    z: p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //+z+y-x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: p.z,
                    y: p.y,
                    z: -p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));
        //-z-y-x
        rotations.push(Vec::from_iter(
            beacons
                .iter()
                .map(|p| Point3D {
                    x: -p.z,
                    y: -p.y,
                    z: -p.x,
                })
                .collect::<Vec<Point3D>>(),
        ));

        Scanner3D {
            rotations,
            matched_rotation_index: 0,
            matched_position: Point3D { x: 0, y: 0, z: 0 },
        }
    }

    pub fn overlap(self: &Self, other: &mut Self) -> bool {
        let base_points = &self.rotations[self.matched_rotation_index];
        let mut distances = Vec::new();
        for r in 0..24 {
            let target_points = &other.rotations[r];
            for base in base_points {
                for target in target_points {
                    distances.push(*base - *target);
                }
            }
            let frequencies = distances
                .iter()
                .copied()
                .fold(HashMap::new(), |mut map, val| {
                    map.entry(val).and_modify(|frq| *frq += 1).or_insert(1);
                    map
                });
            let mut max_p = Point3D { x: 0, y: 0, z: 0 };
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
                for i in 0..other.rotations.len() {
                    for j in 0..other.rotations[i].len() {
                        other.rotations[i][j] = other.rotations[i][j] + max_p;
                    }
                }
                return true;
            }
            distances.clear();
        }
        return false;
    }
}

pub fn run() {
    let input = include_str!("../input/day19");
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
            for j in 0..scanners.len() {
                let overlap = s.overlap(&mut scanners[j]);
                if overlap {
                    aligned.push(scanners.remove(j));
                    break;
                }
            }
        }
    }

    let mut all_beacons = HashSet::new();

    for scanner in &aligned {
        for beacon in &scanner.rotations[scanner.matched_rotation_index] {
            all_beacons.insert(*beacon);
        }
    }

    println!("19.1: {:?}", all_beacons.len());

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

    println!("19.2: {:?}", *m_distances.iter().max().unwrap());
}
