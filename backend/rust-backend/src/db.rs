use diesel::pg::PgConnection;
use diesel::prelude::*;
use dotenv::dotenv;
use std::env;
use serde::{Deserialize, Serialize};
use crate::schema::tasks;
use crate::model;

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");

    PgConnection::establish(&database_url)
        .expect(&format!("Error connection to {}", database_url))
}

#[derive(Serialize, Deserialize, AsChangeset, Insertable, Queryable)]
#[table_name = "tasks"]
pub struct Task {
    pub id: i32,
    pub label: String,
    pub date: String,
    pub done: bool,
}

impl Task {
    pub fn find_all() -> Result<Vec<Self>, Err> {
        let conn = establish_connection();
        let tasks = tasks::table.load::<Task>(&conn);
        Ok(tasks)
    }

    pub fn find(id: i32) -> Result<Self, Err> {
        let conn = establish_connection();
        let task = employees::table.filter(employees::id.eq(id)).first(&conn)?;
        Ok(task)
    }

    pub fn create(task: model::Task) -> Result<Self, Err> {
        let conn = establish_connection();
        let task = Task::from(task);
        let task = diesel::insert_into(tasks::table)
            .values(task)
            .get_result(&conn)?;
        Ok(task)
    }

    pub fn update(id: i32, task: model::Task) -> Result<Self, Err> {
        let conn = establish_connection();
        let task = diesel::update(tasks::table)
            .filter(tasks::id.eq(id))
            .set(task)
            .get_result(&conn)?;
        Ok(task)
    }

    pub fn delete(id: i32) -> Result<usize, Err> {
        let conn = establish_connection();
        let result = diesel::delete(tasks::table.filter(tasks::id.eq(id)))
            .execute(&conn)?;
        Ok(result)
    }
}

impl Task {
    fn from(task: model::Task) -> Task {
      Task {
          id: task.id,
          label: task.label,
          date: task.date,
          done: task.done,
      }
    }
}