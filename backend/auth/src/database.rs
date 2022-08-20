use deadpool_postgres::{ManagerConfig, RecyclingMethod,Client};
use deadpool_postgres::Pool;
use postgres::NoTls;
use std::env::{self};

fn env_parse(name: &str, default: String) -> String{

    match env::var(name){
        Ok(val) => val,
        VarError => default
    }

}

fn get_db_config() -> deadpool_postgres::Config {
    
    let mut config = deadpool_postgres::Config::new();
    config.user = Some(env_parse("DB_USER", "postgres".into()));
    config.password = Some(env_parse("DB_PASSWORD", "Aa123456".into()));
    config.dbname = Some(env_parse("DB_NAME", "restaurant".into()));
    config.host = Some(env_parse("DB_HOSTNAME", "localhost".into()));
    
    config.manager =
       Some(ManagerConfig { recycling_method: RecyclingMethod::Fast });

    config


}



pub fn create_pool() -> Result<Pool, String> {
    Ok(get_db_config().create_pool(NoTls).map_err(|err| err.to_string())?)
}

pub  fn get_connection(pool: &Pool) -> Result<Client, String> {
    pool.get().map_err(|err| err.to_string())
}

pub struct User{
    id: String,
    username: String,
    password: String,
    is_admin: bool,
}