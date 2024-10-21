use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct LoginPayload {
    username: String,
    password: String,
}
#[derive(Debug, Serialize)]
pub struct LoginResponse {
    message: String,
}

pub async fn controller(
    cookies: Cookies,
    Json(payload): Json<LoginPayload>,
) -> Result<Json<LoginResponse>> {
    println!("->> {:<12} - api-login", "Handler");

    if payload.username != "demo" || payload.password != "welcome" {
        return Err(Error::LoginFail);
    }

    // todo: implement real auth token generation
    cookies.add(Cookie::new("auth-token", AUTH_TOKEN));

    let body_response = Json(LoginResponse {
        message: "".to_string(),
    });

    Ok(body_response)
}
// #endregion
