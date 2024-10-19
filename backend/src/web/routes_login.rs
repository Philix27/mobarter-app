use crate::{web::AUTH_TOKEN, Error, Result};

use axum::{routing::post, Json, Router};
use serde::Deserialize;
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct LoginPayload {
    username: String,
    password: String,
}

pub async fn api_login(cookies: Cookies, Json(payload): Json<LoginPayload>) -> Result<Json<Value>> {
    // let cookies: Cookies;

    println!("->> {:<12} - api-login", "Handler");

    if payload.username != "demo" || payload.password != "welcome" {
        return Err(Error::LoginFail);
    }

    // todo: implement real auth token generation
    cookies.add(Cookie::new("auth-token", AUTH_TOKEN));

    let body_response = Json(json!({
        "result": {
            "success": true
        }
    }));

    Ok(body_response)
}
