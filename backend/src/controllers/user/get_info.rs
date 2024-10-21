use crate::Result;
use axum::{routing::post, Json};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct KycBvnPayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycBvnResponse {
    message: String,
}

pub async fn get_info(
    cookies: Cookies,
    Json(payload): Json<KycBvnPayload>,
) -> Result<Json<KycBvnResponse>> {
    let body_response = Json(KycBvnResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
