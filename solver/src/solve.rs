pub fn solve(width: usize, height: usize) -> (Vec<Vec<String>>, f64) {
    let timer = Timer::new();
    let mut board = vec![vec![String::from(""); width]; height];
    let mut pieces = vec![
        Piece {
            color: String::from("#fdf100"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 1, y: 0 },
                Point { x: 2, y: 0 },
                Point { x: 2, y: 1 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#29005d"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 0, y: 2 },
                Point { x: 1, y: 0 },
                Point { x: 1, y: -1 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#66e25a"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 1, y: 0 },
                Point { x: 1, y: -1 },
                Point { x: 2, y: -1 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#BB0000"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 1, y: 0 },
                Point { x: 2, y: 0 },
                Point { x: 1, y: 1 },
                Point { x: 1, y: -1 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#996e5b"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 1, y: 0 },
                Point { x: 1, y: 1 },
                Point { x: 1, y: -1 },
                Point { x: 1, y: -2 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#234c83"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 1, y: 0 },
                Point { x: 1, y: -1 },
                Point { x: 1, y: -2 },
                Point { x: 2, y: -2 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#808080"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 1, y: 0 },
                Point { x: 1, y: -1 },
                Point { x: 1, y: -2 },
                Point { x: 2, y: -1 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#000080"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 0, y: 2 },
                Point { x: 0, y: 3 },
                Point { x: 0, y: 4 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#dad400"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 0, y: 2 },
                Point { x: 0, y: 3 },
                Point { x: 1, y: 0 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#62b7ff"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 0, y: 2 },
                Point { x: 1, y: 0 },
                Point { x: 2, y: 0 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#ffc0cb"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 0, y: 1 },
                Point { x: 0, y: 2 },
                Point { x: 1, y: 1 },
                Point { x: 1, y: 2 },
            ],
            used: false,
        },
        Piece {
            color: String::from("#004900"),
            positions: vec![
                Point { x: 0, y: 0 },
                Point { x: 1, y: 0 },
                Point { x: 2, y: 0 },
                Point { x: 1, y: -1 },
                Point { x: 1, y: -2 },
            ],
            used: false,
        },
    ];

    let mut solutions = Vec::new();

    search(&mut board, &mut pieces, &mut solutions);

    for i in 0..solutions.len() {
        for j in 0..solutions[i].len() {
            if solutions[i][j].eq("") {
                solutions[i][j] = String::from("#000000");
            }
        }
    }

    return (solutions, timer.get_time());
}

fn search(board: &mut Vec<Vec<String>>, pieces: &mut Vec<Piece>, solutions: &mut Vec<Vec<String>>) {
    let (width, height) = (board[0].len(), board.len());
    let mut idx = 0;
    for i in 0..60 {
        let (x, y) = (i / height, i % height);
        if board[y][x].eq("") {
            idx = i;
            break;
        }
    }
    let mut remaining_count = 0;
    for i in 0..pieces.len() {
        if pieces[i].used == false {
            remaining_count += 1;
        }
    }
    for i in 0..pieces.len() {
        if pieces[i].used {
            continue;
        }
        if solutions.len() > 0 {
            break;
        }
        for p in pieces[i].all_angle() {
            if fill(board, idx, p.clone(), false) {
                if remaining_count == 1 {
                    println!("found");
                    let mut new_solution = vec![String::from(""); 60];
                    for i in 0..height {
                        for j in 0..width {
                            new_solution[i * width + j] = board[i][j].clone();
                        }
                    }
                    solutions.push(new_solution);
                    break;
                } else {
                    pieces[i].used = true;
                    search(board, pieces, solutions);
                }
                fill(board, idx, p.clone(), true);
                pieces[i].used = false;
            }
        }
    }
}

fn fill(board: &mut Vec<Vec<String>>, idx: usize, piece: Piece, reset: bool) -> bool {
    let (width, height) = (board[0].len() as i32, board.len() as i32);
    let mut target_points = vec![];
    for &p in piece.positions.iter() {
        let new_target = Point {
            x: idx as i32 / height + p.x,
            y: idx as i32 % height + p.y,
        };
        if new_target.x >= width || new_target.y >= height || new_target.x < 0 || new_target.y < 0 {
            return false;
        }
        if board[new_target.y as usize][new_target.x as usize] != "" && !reset {
            return false;
        }
        target_points.push(new_target);
    }

    for &p in target_points.iter() {
        if reset {
            board[p.y as usize][p.x as usize] = String::from("");
        } else {
            board[p.y as usize][p.x as usize] = piece.color.clone();
        }
    }

    return true;
}

#[derive(Debug, Copy, Clone)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug, Clone)]
struct Piece {
    color: String,
    positions: Vec<Point>,
    used: bool,
}

impl Piece {
    fn rotate90(&self) -> Piece {
        let mut base_position = self.positions[0];
        for &p in self.positions.iter() {
            if p.y > base_position.y {
                base_position = p;
            } else if p.y == base_position.y && p.x < base_position.x {
                base_position = p;
            }
        }

        let mut new_positions = vec![];
        for &p in self.positions.iter() {
            let new_position = Point {
                x: base_position.y - p.y,
                y: p.x - base_position.x,
            };
            new_positions.push(new_position);
        }
        let rotated_pieces = Piece {
            color: self.color.clone(),
            positions: new_positions,
            used: self.used,
        };

        return rotated_pieces;
    }

    fn rotate180(&self) -> Piece {
        let rotate90_piece = self.rotate90();
        return rotate90_piece.rotate90();
    }

    fn rotate270(&self) -> Piece {
        let rotate180_piece = self.rotate180();
        return rotate180_piece.rotate90();
    }

    fn flip_x(&self) -> Piece {
        let mut base_position = self.positions[0];
        for &p in self.positions.iter() {
            if p.x > base_position.x {
                base_position = p;
            } else if p.x == base_position.x && p.y < base_position.y {
                base_position = p;
            }
        }

        let mut new_positions = vec![];
        for &p in self.positions.iter() {
            let new_position = Point {
                x: base_position.x - p.x,
                y: p.y - base_position.y,
            };
            new_positions.push(new_position);
        }
        let fliped_pieces = Piece {
            color: self.color.clone(),
            positions: new_positions,
            used: self.used,
        };

        return fliped_pieces;
    }

    fn flip_y(&self) -> Piece {
        let mut base_position = self.positions[0];
        for &p in self.positions.iter() {
            if p.y > base_position.y {
                base_position = p;
            } else if p.y == base_position.y && p.x < base_position.x {
                base_position = p;
            }
        }

        let mut new_positions = vec![];
        for &p in self.positions.iter() {
            let new_position = Point {
                x: p.x - base_position.x,
                y: base_position.y - p.y,
            };
            new_positions.push(new_position);
        }
        let fliped_pieces = Piece {
            color: self.color.clone(),
            positions: new_positions,
            used: self.used,
        };

        return fliped_pieces;
    }

    fn all_angle(&self) -> Vec<Piece> {
        let mut pieces = vec![
            self.clone(),
            self.rotate90(),
            self.rotate180(),
            self.rotate270(),
        ];

        for i in 0..4 {
            pieces.push(pieces[i].flip_x());
            pieces.push(pieces[i].flip_y());
        }

        return pieces;
    }
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
