use crate::{Error, Result};

// use  anyhow::Ok;
// use anyhow::Ok;
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
pub struct ModelController {
    pub ticket_store: Arc<Mutex<Vec<Option<Ticket>>>>,
}

impl ModelController {
    pub fn new() -> Result<Self> {
        Ok(Self {
            ticket_store: Arc::default(),
        })
    }

    pub fn create_ticket(&self, ticket_fc: TicketForCreate) -> Result<Ticket> {
        let mut store = self.ticket_store.lock().unwrap();

        let id = store.len() as u64;
        let ticket = Ticket {
            id,
            title: ticket_fc.title,
        };

        store.push(Some(ticket.clone()));

        Ok(ticket)
    }

    pub fn list_tickets(&self) -> Result<Vec<Ticket>> {
        let store = self.ticket_store.lock().unwrap();
        let iter = store.iter();

        let tickets = iter.map(|t| t.clone().unwrap()).collect();
        // let tickets = store.iter().filter_map(|t| t.clone());

        Ok(tickets)
    }

    pub fn delete_tickets(&self, id: u64) -> Result<Ticket> {
        let mut store = self.ticket_store.lock().unwrap();

        let ticket = store.get_mut(id as usize).and_then(|t| t.take());

        ticket.ok_or(Error::TicketDeleteNotFound { id: id })
    }
}
