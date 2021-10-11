use serde::{Desrialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Task {
    pub id: i32,
    pub label: String,
    pub date: Date,
    pub done: bool,
}