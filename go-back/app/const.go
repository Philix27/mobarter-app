package app

import "github.com/graphql-go/graphql"

var (
	ID = &graphql.Field{
		Type: graphql.ID}

	DateTime = &graphql.Field{
		Type: graphql.DateTime,
	}

	Boolean = &graphql.Field{
		Type: graphql.Boolean,
	}

	Float = &graphql.Field{
		Type: graphql.Float,
	}

	String = &graphql.Field{
		Type: graphql.String,
	}

	Int = &graphql.Field{
		Type: graphql.Int,
	}
)

var (
	ArgInt = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}
	ArgString = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}
	ArgOptionalString = &graphql.ArgumentConfig{
		Type: graphql.String,
	}
)

const (
	TransactionAirtime   = "Airtime"
	TransactionData      = "Data"
	TransactionTransfer  = "Transfer"
	TransactionCable     = "Cable Tv"
	TransactionLightBill = "Light Bill"
)
