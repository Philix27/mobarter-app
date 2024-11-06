use reqwest::Client;
use reqwest::Error;
use serde::{Deserialize, Serialize};
// use crate::Error;

#[derive(Debug, Deserialize, Serialize)]
struct BitmamaPayload {
    coin: String,
    ticker: String,
    fiatCurrency: String,
    amountOfCryptoToSell: i32,
}

pub async fn buy() -> Result<(), Error> {
    let client = Client::new();

    let payload = BitmamaPayload {
        coin: "usdt".to_string(),
        ticker: "usdngn".to_string(),
        fiatCurrency: "ngn".to_string(),
        amountOfCryptoToSell: 10,
    };
    println!("in bitmama");
    let response = client
        .post("https://walletapp.bitmama.io/buySell/sell/rate?bitmama-oracle-tracker=m315uwd2")
        .json(&payload)
        .send()
        .await?;

    // if response.is_ok() {
    //     println!(
    //         "Request was successful! {:#?}",
    //         response.unwrap().text().unwrap()
    //     );
    // } else {
    //     println!("Failed to send request: {:?}", response.err());
    // }
    // Check the response status
    if response.status().is_success() {
        println!("Request was successful! {:#?}", response);
    } else {
        println!("Failed to send request: {:?}", response.status());
    }

    Ok(())
}
