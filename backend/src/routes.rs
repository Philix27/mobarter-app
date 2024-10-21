use axum::{routing::post, Router};

use crate::{
    app_state::AppState,
    controllers::auth::{self, login},
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
    Router::new().route("/auth/login", post(auth::login))
}
fn airtime_routes(app_state: AppState) -> Router {
    Router::new().route("/airtime", post(login))
}
fn user_routes(app_state: AppState) -> Router {
    Router::new().route("/user", post(login))
}
fn kyc_routes(app_state: AppState) -> Router {
    Router::new().route("/kyc", post(login))
}
fn bank_accounts_routes(app_state: AppState) -> Router {
    Router::new().route("/bank_accounts", post(login))
}
fn support_routes(app_state: AppState) -> Router {
    Router::new().route("/support", post(login))
}
fn orders_routes(app_state: AppState) -> Router {
    Router::new().route("/orders", post(login))
}
fn swap_routes(app_state: AppState) -> Router {
    Router::new().route("/swap", post(login))
}
fn exchange_routes(app_state: AppState) -> Router {
    Router::new().route("/exchange", post(login))
}
fn fiat_walletroutes(app_state: AppState) -> Router {
    Router::new().route("/fiat_wallet", post(login))
}
