use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct SendOtpPayload {
    email: String,
}
#[derive(Debug, Serialize)]
pub struct SendOtpResponse {
    message: String,
}

pub async fn controller(
    cookies: Cookies,
    Json(payload): Json<SendOtpPayload>,
) -> Result<Json<SendOtpResponse>> {
    todo!()
}
// #endregion
