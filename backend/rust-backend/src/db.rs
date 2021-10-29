use diesel::pg::PgConnection;
use diesel::r2d2::ConnectionManager;
use lazy_static::lazy_static;
use r2d2;
use std::env;

use crate::errors::TaskError;

type Pool = r2d2::Pool<ConnectionManager<PgConnection>>;
pub type DBConnection = r2d2::PooledConnection<ConnectionManager<PgConnection>>;


// embed_migrations!();

lazy_static! {
    static ref POOL: Pool = {
        let db_url = env::var("DATABASE_URL").expect("DATABASE_URL variable not set");
        let manager = ConnectionManager::<PgConnection>::new(db_url);
        Pool::new(manager).expect("Failed to creat to db pool")
    };
}

// pub fn init() {
//     // info!("Initializing DB");
//     lazy_static::initialize(&POOL);
//     let conn = connection().expect("Failed to get db connection");
//     embedded_migrations::run(&conn).unwrap();
// }

pub fn connection() -> Result<DBConnection, TaskError> {
    POOL.get()
        .map_err(|e| TaskError::new(500, format!("Failed getting db connection: {}", e)))
}