use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct KycNinPayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycNinResponse {
    message: String,
}

pub async fn approve(
    cookies: Cookies,
    Json(payload): Json<KycNinPayload>,
) -> Result<Json<KycNinResponse>> {
    let body_response = Json(KycNinResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
