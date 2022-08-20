#[macro_use]
extern crate lazy_static;

use std::sync::Arc;

use warp::{self, Filter};
use crate::handlers::*;
use deadpool_postgres::{Pool,Client};
use database::{create_pool, get_connection};

mod handlers;
mod auth;
mod database;


lazy_static!{
    pub static ref pool: Arc<Pool> = Arc::new(match create_pool(){
        Ok(p) => p,
        Err(err) => panic!("{}", err),
    });
}




#[tokio::main]
async fn main() {

    let a : Result<Client,String> = get_connection(&pool);
    let client = a.unwrap();

    

    

    let login_route = warp::path!("login")
    .and(warp::post())
   .and(warp::body::json())
    .and_then(login_handler);
    let user_route = warp::path!("user")
    .and_then(user_handler);

    let admin_route = warp::path!("admin")
    .and_then(admin_handler);

    let routes = login_route
    .or(user_route)
    .or(admin_route);

    warp::serve(routes).run(([0,0,0,0],8080)).await;

}   




