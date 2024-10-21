mod app_state;
mod controllers;
mod ctx;
mod error;
mod model;
mod mw;
mod routes;
mod web;

pub use self::error::{Error, Result};

use app_state::AppState;
use axum::{
    extract::{Path, Query},
    middleware,
    response::{IntoResponse, Response},
    routing::get,
    Router,
};

use model::ModelController;
use routes::core_routes;
use serde::Deserialize;
use tower_cookies::CookieManagerLayer;
use tower_http::services::ServeDir;

#[shuttle_runtime::main]
async fn main() -> shuttle_axum::ShuttleAxum {
    // ServeDir falls back to serve index.html when requesting a directory
    // let router = Router::new().nest_service("/", ServeDir::new("assets"));

    // Ok(router.into())

    // build our application with a single route

    let mc = ModelController::new().unwrap();
    let app_state = AppState::new().unwrap();
    let routes_apis =
        web::tickets::routes(mc).route_layer(middleware::from_fn(web::mw_auth::mw_require_auth));
    let core_auth_routes =
        core_routes(app_state).route_layer(middleware::from_fn(web::mw_auth::mw_require_auth));

    let app = Router::new()
        .merge(static_routes())
        .merge(web::routes())
        .nest("/api", routes_apis)
        .nest("/api/v1", core_auth_routes)
        .layer(middleware::map_response(main_response_mapper))
        .layer(CookieManagerLayer::new())
        .route("/hey", get(|| async { "Hello, World!" }))
        .route("/foo", get(get_foo).post(post_foo))
        .route("/query_param", get(handle_hello))
        .route("/path_param/:name", get(handle_path))
        .fallback_service(get(handle_hello));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:5000").await.unwrap();
    axum::serve(listener, app.clone()).await.unwrap();

    Ok(app.into())
}

async fn main_response_mapper(res: Response) -> Response {
    println!(" ->> RESPONSE MW ");
    println!();
    res
}

#[derive(Debug, Deserialize)]
struct HelloParams {
    name: Option<String>,
}
// which calls one of these handlers
async fn handle_hello(Query(params): Query<HelloParams>) -> impl IntoResponse {
    println!("->> {:<12} - handle_hello - {params:?}", "Handler");

    let name = params.name.as_deref().unwrap_or("World");
    println!("->> Query Params - {}", name);
}
// which calls one of these handlers
async fn handle_path(Path(params): Path<HelloParams>) -> impl IntoResponse {
    println!("->> {:<12} - handle_path - {params:?}", "Handler");

    let name = params.name.as_deref().unwrap_or("World");
    println!("->> Path Params - {}", name);
}

async fn get_foo() {}
async fn post_foo() {}
async fn foo_bar() {}

// Router

fn static_routes() -> Router {
    Router::new().nest_service("/", ServeDir::new("./"))
}

// async fn getService() {}
