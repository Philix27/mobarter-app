use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct KycPhonePayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycPhoneResponse {
    message: String,
}

pub async fn verify_phone(
    cookies: Cookies,
    Json(payload): Json<KycPhonePayload>,
) -> Result<Json<KycPhoneResponse>> {
    let body_response = Json(KycPhoneResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
