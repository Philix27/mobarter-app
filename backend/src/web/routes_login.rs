use crate::{web::AUTH_TOKEN, Error, Result};

use axum::{routing::post, Json, Router};
use serde::Deserialize;
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

pub fn routes() -> Router {
    Router::new().route("/auth/login", post(api_login))
}

#[derive(Debug, Deserialize)]
struct LoginPayload {
    username: String,
    password: String,
}

fn api_login(cookies: Cookies, Json(payload): Json<LoginPayload>) -> Result<Json<Value>> {
    println!("->> {:<12} - api-login", "Handler");

    if payload.username != "demo" || payload.password != "welcome" {
        return Err(Error::LoginFail);
    }

    // todo: implement real cookiw
    cookies.add(Cookie::new("auth-token", AUTH_TOKEN));

    let body_response = Json(json!({
        "result": {
            "success": true
        }
    }));

    Ok(body_response)
}
