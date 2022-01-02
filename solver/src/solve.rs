pub fn solve() -> (Vec<Vec<String>>, f64) {
    let timer = Timer::new();

    let mut sol_tmp = vec![String::from("#AA0000"); 30];

    for _ in 0..30 {
        sol_tmp.push(String::from("#00AA00"));
    }

    let sol = vec![sol_tmp.clone(); 4];

    return (sol, 2.0);
}

struct Piece {
    color: String,
    position: Vec<(i32, i32)>,
}

fn get_time() -> f64 {
    let t = std::time::SystemTime::now()
        .duration_since(std::time::UNIX_EPOCH)
        .unwrap();
    t.as_secs() as f64 + t.subsec_nanos() as f64 * 1e-9
}

struct Timer {
    start_time: f64,
}

impl Timer {
    fn new() -> Timer {
        Timer {
            start_time: get_time(),
        }
    }

    fn get_time(&self) -> f64 {
        get_time() - self.start_time
    }

    #[allow(dead_code)]
    fn reset(&mut self) {
        self.start_time = 0.0;
    }
}
