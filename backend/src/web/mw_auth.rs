use axum::{extract::Request, middleware::Next, response::Response};
use lazy_regex::regex_captures;
use tower_cookies::Cookies;

use crate::{ctx::Ctx, web::AUTH_TOKEN, Error, Result};

pub async fn mw_require_auth(ctx: Result<Ctx>, req: Request, next: Next) -> Result<Response> {
    ctx?;
    Ok(next.run(req).await)
}
