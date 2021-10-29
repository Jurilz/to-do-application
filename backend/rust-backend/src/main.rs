#[macro_use]
extern crate actix_web;
#[macro_use]
extern crate diesel;
#[macro_use]
extern crate serde_json;

use std::env;
use actix_web::{App, HttpServer};

mod db;
mod schema;
mod errors;
mod task;


#[actix_rt::main]
async fn main() -> std::io::Result<()>{

    dotenv::dotenv().ok();
    env_logger::init();

    let host = env::var("HOST").expect("HOST variable not set");
    let port = env::var("PORT").expect("PORT variable not set");

    HttpServer::new(move || {
        App::new()
            .configure(task::init_routes)
    })
        .bind(format!("{}:{}", host, port))?
        .run()
        .await
}
