use std::fs;

fn a (ddims: &Vec<Dims>) -> i32 {
    let mut total_paper: i32 = 0;
    for dims in ddims {
        total_paper += dims.paper_needed();
    }

    return total_paper;
}

fn b (ddims: &Vec<Dims>) -> i32 {
    let mut total_ribbon: i32 = 0;
    for dims in ddims {
        total_ribbon += dims.ribbon_needed();
    }

    return total_ribbon;
}

#[derive(Debug)]
struct Dims {
    l: i32,
    w: i32,
    h: i32,
}

impl Dims {
    fn surface_area(&self) -> i32 {
        return (2*self.l*self.w) + (2*self.w*self.h) + (2*self.h*self.l);
    }

    fn paper_needed(&self) -> i32 {
        let areas = [self.l*self.w, self.w*self.h, self.l*self.h];
        return *areas.iter().min().expect("Couldn't get smallest area") + self.surface_area();
    }

    fn volume(&self) -> i32 {
        return self.l*self.w*self.h;
    }

    fn min_perimeter(&self) -> i32 {
        let perims = [
            (self.l*2) + (self.w*2),
            (self.w*2) + (self.h*2),
            (self.l*2) + (self.h*2),
        ];

        return *perims.iter().min().expect("Couldn't get smallest perimeter");
    }

    fn ribbon_needed(&self) -> i32 {
        return self.min_perimeter() + self.volume()
    }
}

fn parse_dims(s: &str) -> Dims {
    let nums = s.split("x")
        .map(|n|
            n.parse::<i32>().expect(&format!("Could not parse num: {}", n)))
        .collect::<Vec<_>>();

    return Dims {
        l: nums[0],
        w: nums[1],
        h: nums[2]
    };
}

y2015::main! {
    let ds = &fs::read_to_string("./input/day02.txt")
        .expect("Could not open input")
        .lines()
        .map(|line| parse_dims(line))
        .collect::<Vec<_>>();

    (a(ds), b(ds))
}
