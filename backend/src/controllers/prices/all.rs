use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct KycStatusPayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycStatusResponse {
    message: String,
}

pub async fn all(
    cookies: Cookies,
    Json(payload): Json<KycStatusPayload>,
) -> Result<Json<KycStatusResponse>> {
    let body_response = Json(KycStatusResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
