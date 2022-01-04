use actix_cors::Cors;
use actix_web::{get, web, App, HttpResponse, HttpServer, http};
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

#[get("/solve/{height}/{width}")]
async fn index(size: web::Path<Size>) -> HttpResponse {
    let sol = solve::solve(size.height, size.width);
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
    HttpServer::new(|| {
        let cors = Cors::default()
            .allowed_origin_fn(|_, _| true)
            .allowed_methods(vec!["GET"])
            .allowed_headers(vec![http::header::AUTHORIZATION, http::header::ACCEPT])
            .allowed_header(http::header::CONTENT_TYPE)
            .max_age(3600);

        App::new().wrap(cors).service(index)
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}
