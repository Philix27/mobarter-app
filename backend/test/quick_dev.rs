#![allow(unused)]

use anyhow::Result;
use httpc_test;

#[tokio::test]
async fn quick_dev() -> Result<()> {
    let hc = httpc_test::new_client("http://localhost:8080")?;

    hc.do_get("/hey").await?.print().await?;

    Ok(())
}
