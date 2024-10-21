use axum::{routing::post, Router};

use crate::{controllers::auth::login, model::ModelController};

pub fn auth_route(mc: ModelController) -> Router {
    Router::new().route("/auth/login", post(login))
}
pub fn airtime_routes(mc: ModelController) -> Router {
    Router::new().route("/airtime", post(login))
}
pub fn user_routes(mc: ModelController) -> Router {
    Router::new().route("/user", post(login))
}
pub fn kyc_routes(mc: ModelController) -> Router {
    Router::new().route("/kyc", post(login))
}
pub fn bank_accounts_routes(mc: ModelController) -> Router {
    Router::new().route("/bank_accounts", post(login))
}
pub fn support_routes(mc: ModelController) -> Router {
    Router::new().route("/support", post(login))
}
pub fn orders_routes(mc: ModelController) -> Router {
    Router::new().route("/orders", post(login))
}
pub fn swap_routes(mc: ModelController) -> Router {
    Router::new().route("/swap", post(login))
}
pub fn exchange_routes(mc: ModelController) -> Router {
    Router::new().route("/exchange", post(login))
}
pub fn fiat_walletroutes(mc: ModelController) -> Router {
    Router::new().route("/fiat_wallet", post(login))
}
