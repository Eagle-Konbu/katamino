use actix_web::{get, web, App, HttpResponse, HttpServer};
use serde::{Deserialize, Serialize};

mod solve;

#[derive(Deserialize)]
struct Size {
    width: usize,
    height: usize,
}

#[derive(Serialize, Clone)]
struct Solution {
    width: usize,
    height: usize,
    calc_time: f64,
    solutions: Vec<Vec<String>>,
}

#[get("/solve/{width}/{height}")]
async fn index(size: web::Path<Size>) -> HttpResponse {
    let sol = solve::solve();
    let res = Solution {
        width: size.width,
        height: size.height,
        calc_time: sol.1,
        solutions: sol.0,
    };

    HttpResponse::Ok().json(res)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(index))
        .bind("0.0.0.0:8080")?
        .run()
        .await
}
