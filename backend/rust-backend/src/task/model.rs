use diesel::prelude::*;
use serde::{Deserialize, Serialize};

use crate::errors::TaskError;
use crate::db;
use crate::schema::{tasks};

#[derive(Serialize, Deserialize, Queryable, Identifiable, Insertable, AsChangeset)]
#[table_name = "tasks"]
pub struct Task {
    pub id: i32,
    pub label: String,
    pub date: String,
    pub done: bool
}

impl Task {

    pub fn get_all() -> Result<Vec<Self>, TaskError> {
        let conn = db::connection()?;
        let tasks = tasks::table.load::<Task>(&conn)?;
        Ok(tasks)
    }

    pub fn get(id:i32) -> Result<Task, TaskError> {
        let conn = db::connection()?;
        let task = tasks::table.filter(tasks::id.eq(id)).first(&conn)?;
        Ok(task)
    }

    pub fn create(task: Task) -> Result<Self, TaskError> {
        let conn = db::connection()?;
        let task = diesel::insert_into(tasks::table)
            .values(task)
            .get_result(&conn)?;
        Ok(task)
    }

    pub fn update(id: i32, task: Task) -> Result<Self, TaskError> {
        let conn = db::connection()?;
        let task = diesel::update(tasks::table)
            .filter(tasks::id.eq(id))
            .set(task)
            .get_result(&conn)?;
        Ok(task)
    }

    pub fn delete(id: i32) -> Result<usize, TaskError> {
        let conn = db::connection()?;
        let result = diesel::delete(tasks::table.filter(tasks::id.eq(id)))
            .execute(&conn)?;
        Ok(result)
    }

}