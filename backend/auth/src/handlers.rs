use chrono::Utc;
use serde::{Deserialize, Serialize};
use warp::{Reply, Rejection, reply};
use crate::auth::*;
type WebResult<T> = std::result::Result<T, Rejection>;

#[derive(Deserialize)]
pub struct LoginCredentials{
    pub username: String,
    pub password: String
}

#[derive(Serialize)]
pub struct LoginResponse {
    pub token: String,
}



pub async fn login_handler(body: LoginCredentials) ->  WebResult<impl Reply> {
        println!("username: {}, password: {}",body.username, body.password);

        // add authentication


        let token = create_token(body.username,&Role::User);


        Ok(reply::json(&LoginResponse { token }))
}

pub async fn user_handler() -> WebResult<impl Reply> {
    Ok("User")
}

pub async fn admin_handler() -> WebResult<impl Reply> {
    Ok("Admin")
}