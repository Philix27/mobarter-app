use axum::{extract::FromRequestParts, http::request::Parts, RequestPartsExt};
use lazy_regex::regex_captures;
use tower_cookies::Cookies;

use crate::{web::AUTH_TOKEN, Error, Result};

#[derive(Debug, Clone)]
pub struct Ctx {
    user_id: u64,
}

impl Ctx {
    pub fn new(user_id: u64) -> Self {
        Self { user_id }
    }
}

impl Ctx {
    pub fn user_id(&self) -> u64 {
        self.user_id
    }
}

impl<S: Send + Sync> FromRequestParts<S> for Ctx {
    type Rejection = Error;

    #[doc = " Perform the extraction."]
    #[must_use]
    #[allow(
        elided_named_lifetimes,
        clippy::type_complexity,
        clippy::type_repetition_in_bounds
    )]
    fn from_request_parts<'life0, 'life1, 'async_trait>(
        parts: &'life0 mut Parts,
        state: &'life1 S,
    ) -> ::core::pin::Pin<
        Box<
            dyn ::core::future::Future<Output = Result<Self>> + ::core::marker::Send + 'async_trait,
        >,
    >
    where
        'life0: 'async_trait,
        'life1: 'async_trait,
        Self: 'async_trait,
    {
        todo!()
    }

    // async fn from_request_parts(parts: &mut Parts, _state: &S) -> Result<Self> {
    //     let cookies = parts.extract::<Cookies>().await.unwrap();

    //     let auth_token = cookies.get(AUTH_TOKEN).map(|c| c.value().to_string());
    //     let token: ParseToken = auth_token
    //         .ok_or(Error::AuthFailNoAuthTokenCookie)
    //         .and_then(ParseToken::new)?;

    //     Ok(Ctx::new(token.user_id))
    // }

    // async fn from_request_parts<'life0, 'life1, 'async_trait>(
    //     parts: &'life0 mut Parts,
    //     _state: &'life1 S,
    // ) -> Result<Self>
    // where
    //     'life0: 'async_trait,
    //     'life1: 'async_trait,
    //     Self: 'async_trait,
    // {
    //     let cookies = parts.extract::<Cookies>().await.unwrap();

    //     let auth_token = cookies.get(AUTH_TOKEN).map(|c| c.value().to_string());
    //     let token: ParseToken = auth_token
    //         .ok_or(Error::AuthFailNoAuthTokenCookie)
    //         .and_then(ParseToken::new)?;

    //     Ok(Ctx::new(token.user_id))
    // }
}

pub struct ParseToken {
    pub user_id: u64,
    pub exp: String,
    pub sign: String,
}
impl ParseToken {
    pub fn new(token: String) -> Result<Self> {
        let (_whole, user_id, exp, sign) = regex_captures!(r#"^user-(\d+)\.(.+)\.(.+)"#, &token)
            .ok_or(Error::AuthFailTokenWromgFormat)?;

        let user_id: u64 = user_id
            .parse()
            .map_err(|_| Error::AuthFailTokenWromgFormat)?;

        Ok(Self {
            user_id,
            exp: exp.to_string(),
            sign: sign.to_string(),
        })
    }
}
