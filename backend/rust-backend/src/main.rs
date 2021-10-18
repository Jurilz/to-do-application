mod model;
mod db;
mod schema;

use actix_web::{App, HttpRequest, HttpServer, Responder, web};
use actix_web::web::route;
use crate::model::route;

#[actix_web::main]
async fn main() -> std::io::Result<()>{

    HttpServer::new(|| {
        App::new()
            .configure(route::init_routes)
            // .route("/", web::get().to(welcome))
            // .route("/{name}", web::get().to(welcome))
    })
        .bind("127.0.0.1:8000")?
        .run()
        .await
}

async fn welcome(request: HttpRequest) -> impl Responder {
    let name = request
        .match_info()
        .get("name")
        .unwrap_or("World");
    format!("Hello {}!", &name)
}
