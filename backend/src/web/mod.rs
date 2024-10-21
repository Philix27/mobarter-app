use axum::{routing::post, Router};
use routes_login::api_login;


pub mod routes_login;
pub mod tickets;
pub mod mw_auth;

pub const AUTH_TOKEN: &str = "user-id.exp.sign";

pub fn routes() -> Router {
    Router::new().route("/auth/login", post(api_login))
}
