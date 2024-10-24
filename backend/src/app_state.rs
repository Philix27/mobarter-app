use crate::{Error, Result};
use serde::{Deserialize, Serialize};
use std::sync::{Arc, Mutex};

// region: ticket
#[derive(Debug, Clone, Serialize)]
pub struct Ticket {
    pub id: u64,
    pub title: String,
}

#[derive(Debug, Clone, Deserialize)]
pub struct TicketForCreate {
    pub title: String,
}
#[derive(Debug, Clone)]
pub struct AppState {
    pub ticket_store: Arc<Mutex<Vec<Option<Ticket>>>>,
    pub redis_store: Option<String>,
    pub main_db: Option<String>,
}

impl AppState {
    pub fn new() -> Result<Self> {
        Ok(Self {
            ticket_store: Arc::default(),
            redis_store: None,
            main_db: None,
        })
    }
}
