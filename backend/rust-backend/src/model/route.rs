use actix_web::{get, put, post, delete, web, HttpResponse, Responder};
use serde_json::json;
use chrono;
use crate::model::task::Task;

#[get("/tasks")]
async fn find_all() -> impl Responder {
    HttpResponse::Ok().json(vec![
        Task {
            id: 1,
            label: String::from("Rust API All Items"),
            date: chrono::offset::Local::now().to_string(),
            done: false
        }
    ])
}

#[get("/tasks/{id}")]
async fn find() -> impl Responder {
    HttpResponse::Ok().json(Task {
        id: 1,
        label: String::from("Rust API"),
        date: chrono::offset::Local::now().to_string(),
        done: false
    })
}

pub fn init_routes(config: &mut web::ServiceConfig) {
    config.service(find_all);
    config.service(find);
}