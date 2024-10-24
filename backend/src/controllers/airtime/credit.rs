use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};
use utoipa_axum::{router::OpenApiRouter, routes, PathItemExt};

#[derive(Debug, Deserialize, Serialize)]
pub struct AirtimePayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycApproveResponse {
    message: String,
}

pub async fn credit_no(
    cookies: Cookies,
    Json(payload): Json<AirtimePayload>,
) -> Result<Json<KycApproveResponse>> {
    let body_response = Json(KycApproveResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
