use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct LogoutPayload {
    email: String,
}
#[derive(Debug, Serialize)]
pub struct LogoutResponse {
    message: String,
}

pub async fn controller(
    cookies: Cookies,
    Json(payload): Json<LogoutPayload>,
) -> Result<Json<LogoutResponse>> {
    todo!()
}
// #endregion
