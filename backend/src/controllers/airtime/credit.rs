use crate::Result;
use axum::Json;
use serde::{Deserialize, Serialize};
use tower_cookies::Cookies;

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
