use core::fmt;

use chrono::Utc;
use jsonwebtoken::{Algorithm, Header, encode, EncodingKey};
use serde::{Serialize, Deserialize};

const JWT_SECRET: &[u8] = b"secret";

#[derive(Clone, PartialEq)]
pub enum Role {
    User,
    Admin,
}

impl Role {
    pub fn from_string(role: &str) -> Role{
        match role {
            "Admin" => Role::Admin,
            _ => Role::User
        }
    }
}

impl fmt::Display for Role{
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self{
            Role::User => write!(f,"User"),
            Role::Admin => write!(f,"Admin"),
        }
    }

}


#[derive(Debug, Serialize, Deserialize)]
pub struct Claims {
    pub sub: String,
    pub role: String,
    pub exp: usize,
}

pub fn create_token(username: String, role: &Role) -> String{

    let exp = Utc::now()
    .checked_add_signed(chrono::Duration::days(2))
    .expect("valid timestamp")
    .timestamp();   

    let header = Header::new(Algorithm::HS512);
    let claims = Claims{ sub: username, role: role.to_string(), exp:  exp as usize};

    match encode(&header, &claims, &EncodingKey::from_secret(JWT_SECRET)){
        Ok(token) => return token,
        Err(err) => {println!("{}",err);return "".to_string();}
    }
    

}