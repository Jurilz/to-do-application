use actix_web::{delete, get, post, web, HttpResponse};

use crate::task::Task;
use crate::errors::TaskError;

#[get("/tasks/{id}")]
async fn get(id: web::Path<i32>) -> Result<HttpResponse, TaskError> {
    let task = Task::get(id.into_inner())?;
    Ok(HttpResponse::Ok().json(task))
}

#[get("/tasks")]
async fn get_all() -> Result<HttpResponse, TaskError> {
    let tasks = Task::get_all()?;
    Ok(HttpResponse::Ok().json(tasks))
}

#[post("/tasks")]
async fn create(task: web::Json<Task>) -> Result<HttpResponse, TaskError> {
    let task = Task::create(task.into_inner())?;
    Ok(HttpResponse::Created().json(task))
}


#[put("/tasks/{id}")]
async fn update(
    id: web::Path<i32>,
    task: web::Json<Task>
) -> Result<HttpResponse, TaskError> {
    let id = id.into_inner();
    let task = Task::update(id, task.into_inner())?;
    Ok(HttpResponse::Ok().json(task))
}

#[delete("/tasks/{id}")]
async fn delete(id: web::Path<i32>) -> Result<HttpResponse, TaskError> {
    let count = Task::delete(id.into_inner())?;
    Ok(HttpResponse::Ok().json(count))
}

pub fn init_routes(config: &mut web::ServiceConfig) {
    config.service(get);
    config.service(get_all);
    config.service(create);
    config.service(update);
    config.service(delete);
}