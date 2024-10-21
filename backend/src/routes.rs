use axum::{
    routing::{get, post},
    Router,
};

use crate::{
    app_state::AppState,
    controllers::{auth, kyc},
};

pub fn core_routes(app_state: AppState) -> Router {
    Router::new()
        .merge(auth_route(app_state.clone()))
        .merge(airtime_routes(app_state.clone()))
        .merge(user_routes(app_state.clone()))
        .merge(kyc_routes(app_state.clone()))
        .merge(bank_accounts_routes(app_state.clone()))
        .merge(support_routes(app_state.clone()))
        .merge(orders_routes(app_state.clone()))
        .merge(swap_routes(app_state.clone()))
        .merge(exchange_routes(app_state.clone()))
        .merge(fiat_walletroutes(app_state.clone()))
}

fn auth_route(app_state: AppState) -> Router {
    Router::new()
        .route("/auth/login", post(auth::login::controller))
        .route("/auth/logout", post(auth::logout::controller))
        .route("/auth/send_otp", post(auth::send_otp::controller))
        .route("/auth/verify_otp", post(auth::verify_otp::controller))
}
fn airtime_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/airtime/credit", post(auth::login::controller))
        .route("/airtime/history", get(auth::logout::controller))
}
fn user_routes(app_state: AppState) -> Router {
    Router::new().route("/user/info", get(auth::login::controller))
}
fn kyc_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/kyc/phone", post(kyc::verify_phone))
        .route("/kyc/nin", post(kyc::verify_nin))
        .route("/kyc/bvn", post(kyc::verify_bvn))
        .route("/kyc/approve", post(kyc::approve))
        .route("/kyc/status", get(kyc::get_status))
}
fn bank_accounts_routes(app_state: AppState) -> Router {
    Router::new().route(
        "/bank_accounts",
        post(auth::login::controller).get(auth::login::controller),
    )
}
fn exchange_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/exchange/rates", get(auth::login::controller))
        .route("/exchange/partner_rates", get(auth::login::controller))
}
fn fiat_walletroutes(app_state: AppState) -> Router {
    Router::new().route("/fiat_wallet", post(auth::login::controller))
}
fn orders_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/orders", get(auth::login::controller))
        .route("/orders/buy", post(auth::login::controller))
        .route("/orders/sell", post(auth::login::controller))
        .route("/orders/approve", post(auth::login::controller))
        .route("/orders/close", post(auth::login::controller))
}
fn support_routes(app_state: AppState) -> Router {
    Router::new().route("/support", post(auth::login::controller))
}

fn swap_routes(app_state: AppState) -> Router {
    Router::new().route("/swap", post(auth::login::controller))
}
