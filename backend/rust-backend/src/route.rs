use actix_web::{get, put, post, delete, web, HttpResponse, Responder};
use serde_json::json;
use crate::note::Note;
use crate::domainmodel::task;

#[get("/tasks")]
async fn find_all() -> impl Responder {
    HttpResponse::Ok().json(vec![
        Task {
            
        }
    ])
}