use axum::{routing::post, Router};
use routes_login::api_login;

pub mod routes_login;

pub const AUTH_TOKEN: &str = "user-1.wxp.sign";

pub fn routes() -> Router {
    Router::new().route("/auth/login", post(api_login))
}
