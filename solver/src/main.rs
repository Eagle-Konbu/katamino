use actix_web::{get, web, App, HttpServer, HttpResponse};
use serde::{Serialize, Deserialize};

#[derive(Deserialize)]
struct Size {
    width: usize,
    height: usize,
}

#[derive(Serialize)]
struct Solution {
    width: usize,
    height: usize,
    calc_time: f64,
    solutions: Vec<String>,
}

#[get("/{width}/{height}")]
async fn index(size: web::Path<Size>) -> HttpResponse {
    let res = Solution {
        width: size.width,
        height: size.height,
        calc_time: 5.0,
        solutions: vec![String::from("#AA0000"); size.width * size.height],
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
