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

	Str = &graphql.Field{
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
	ArgStr = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}
)
