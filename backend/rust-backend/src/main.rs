mod route;


use actix_web::{web, App, HttpRequest, HttpServer, Responder};

#[actix_web::main]
async fn main() -> std::io::Result<()>{

    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(welcome))
            .route("/{name}", web::get().to(welcome))
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
