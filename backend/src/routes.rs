use axum::{
    routing::{get, post},
    Router,
};

use crate::{
    app_state::AppState,
    controllers::{airtime, auth, bank_accounts, exchange, kyc, orders, prices, user},
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
        .merge(prices_routes(app_state.clone()))
        .merge(fiat_wallet_routes(app_state.clone()))
}

fn auth_route(app_state: AppState) -> Router {
    Router::new()
        .route("/auth/login", post(auth::login::controller))
        .route("/auth/logout", post(auth::logout::controller))
        .route("/auth/send_otp", post(auth::send_otp::controller))
        .route("/auth/verify_otp", post(auth::verify_otp::controller))
        .with_state(app_state)
}
fn airtime_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/airtime/credit", post(airtime::credit_no))
        .route("/airtime/history", get(airtime::get_history))
        .with_state(app_state)
}
fn user_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/user/info", get(user::get_info))
        .with_state(app_state)
}
fn kyc_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/kyc/phone", post(kyc::verify_phone))
        .route("/kyc/nin", post(kyc::verify_nin))
        .route("/kyc/bvn", post(kyc::verify_bvn))
        .route("/kyc/approve", post(kyc::approve))
        .route("/kyc/status", get(kyc::get_status))
        .with_state(app_state)
}
fn bank_accounts_routes(app_state: AppState) -> Router {
    Router::new()
        .route(
            "/bank_accounts",
            post(bank_accounts::add_account).get(bank_accounts::get_account),
        )
        .with_state(app_state)
}
fn exchange_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/exchange/rates", get(exchange::get_rates))
        .route("/exchange/partner_rates", get(exchange::get_partner_rates))
        .with_state(app_state)
}
fn fiat_wallet_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/fiat_wallet", post(auth::login::controller))
        .with_state(app_state)
}
fn orders_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/orders", get(orders::get_all))
        .route("/orders/buy", post(orders::buy))
        .route("/orders/sell", post(orders::sell))
        .route("/orders/approve", post(orders::approve))
        .route("/orders/close", post(orders::close))
        .with_state(app_state)
}

fn prices_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/prices/buy", get(prices::buy))
        .route("/prices/sell", get(prices::sell))
        .with_state(app_state)
}

fn support_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/support", post(auth::login::controller))
        .with_state(app_state)
}

fn swap_routes(app_state: AppState) -> Router {
    Router::new()
        .route("/swap", post(auth::login::controller))
        .with_state(app_state)
}
