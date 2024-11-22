package app

import (
	"errors"

	"github.com/LNMMusic/optional"
	"github.com/graphql-go/graphql"
)

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

// Arg Input
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

// Input
var (
	InputString = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String), // Optional field
	}
	InputStringOptional = &graphql.InputObjectFieldConfig{
		Type: graphql.String, // Optional field
	}
	InputInt = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.Int), // Optional field
	}
	InputIntOptional = &graphql.InputObjectFieldConfig{
		Type: graphql.Int, // Optional field
	}
	InputFloat = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.Float), // Optional field
	}
	InputFloatOptional = &graphql.InputObjectFieldConfig{
		Type: graphql.Float, // Optional field
	}
	InputBoolean = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.Boolean), // Optional field
	}
	InputBooleanOptional = &graphql.InputObjectFieldConfig{
		Type: graphql.Boolean, // Optional field
	}
	InputId = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.ID), // Optional field
	}
	InputIdOptional = &graphql.InputObjectFieldConfig{
		Type: graphql.ID, // Optional field
	}
)

func SetInput(name string, fields *map[string]*graphql.InputObjectFieldConfig) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:   name,
				Fields: fields,
			},
		),
	}
}

func ValidateArg(argName optional.Option[string], p graphql.ResolveParams) (map[string]interface{}, error) {

	filterArg, filterProvided := p.Args[getInput(argName)].(map[string]interface{})
	if !filterProvided {

		return nil, errors.New("")

	}

	return filterArg, nil
}

func getInput(i optional.Option[string]) string {
	if i.IsSome() {
		return i.Unwrap()
	} else {
		return "input"
	}
}
