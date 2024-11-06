use crate::{web::AUTH_TOKEN, Error, Result};
use axum::{routing::post, Json, Router};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use tower_cookies::{Cookie, Cookies};

use super::comapanies::bitmama;

#[derive(Debug, Serialize)]
pub struct BuyResponse {
    company: String,
    fiat: String,
    coin: String,
    rate: i32,
}

#[derive(Debug, Deserialize)]
pub struct KycApprovePayload {
    payload: String,
}
#[derive(Debug, Serialize)]
pub struct KycApproveResponse {
    message: String,
}

pub async fn buy(
    cookies: Cookies,
    Json(payload): Json<KycApprovePayload>,
) -> Result<Json<KycApproveResponse>> {
    bitmama::buy().await;
    print!("Hey guys");
    let body_response = Json(KycApproveResponse {
        message: payload.payload,
    });

    Ok(body_response)
}
