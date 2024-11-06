pub mod bitmama;

// pub use bitmama;
// pub use bitmama::*;

pub enum PriceType {
    Buy,
    Sell,
}

pub struct Prices {
    pub rate: i32,
    pub coin: String,
    pub fiat: String,
}
