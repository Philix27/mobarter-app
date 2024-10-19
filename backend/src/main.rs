mod error;
mod web;
mod model;

pub use self::error::{Error, Result};
use std::ops::Deref;

use axum::{
    extract::{Path, Query},
    middleware,
    response::{IntoResponse, Response},
    routing::get,
    Router,
};

use serde::Deserialize;
use tower_cookies::CookieManagerLayer;
use tower_http::services::ServeDir;

#[shuttle_runtime::main]
async fn main() -> shuttle_axum::ShuttleAxum {
    // ServeDir falls back to serve index.html when requesting a directory
    // let router = Router::new().nest_service("/", ServeDir::new("assets"));

    // Ok(router.into())

    // build our application with a single route
    let app = Router::new()
        .merge(static_routes())
        .merge(web::routes() )
        .layer(middleware::map_response(main_response_mapper))
        .layer(CookieManagerLayer::new())
        // .nest_service("/", ServeDir::new("assets"))
        .route("/hey", get(|| async { "Hello, World!" }))
        .route("/foo", get(get_foo).post(post_foo))
        .route("/query_param", get(handle_hello))
        .route("/path_param/:name", get(handle_path));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:5000").await.unwrap();
    axum::serve(listener, app.clone()).await.unwrap();

    Ok(app.into())
}

async fn main_response_mapper(res: Response) -> Response {
    println!(" ->> RESPONSE MW ",);
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
