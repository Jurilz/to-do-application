use actix_web::{get, put, post, delete, web, HttpResponse, Responder};
use serde_json::json;
use chrono;
use crate::db::Tasks;
use crate::model::task::Task;

#[get("/tasks")]
async fn find_all() -> Result<HttpResponse, Err> {
    let tasks = Tasks::find_all()?;
    Ok(HttpResponse::Ok().json(tasks))
}

#[get("/tasks/{id}")]
async fn find(id: web::Path<i32>) -> Result<HttpResponse, Err> {
    let task = Tasks::find(id.into_inner())?;
    Ok(HttpResponse::Ok().json(task))
}

#[post]

pub fn init_routes(config: &mut web::ServiceConfig) {
    config.service(find_all);
    config.service(find);
}