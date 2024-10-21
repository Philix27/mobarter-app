use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct VerifyOtpPayload {
    email: String,
}
#[derive(Debug, Serialize)]
pub struct VerifyOtpResponse {
    message: String,
}

pub async fn controller(
    cookies: Cookies,
    Json(payload): Json<VerifyOtpPayload>,
) -> Result<Json<VerifyOtpResponse>> {
    todo!()
}
// #endregion
