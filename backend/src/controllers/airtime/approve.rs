use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

#[derive(Debug, Deserialize)]
pub struct KycApprovePayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycApproveResponse {
    message: String,
}

pub async fn approve(
    cookies: Cookies,
    Json(payload): Json<KycApprovePayload>,
) -> Result<Json<KycApproveResponse>> {
    let body_response = Json(KycApproveResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
