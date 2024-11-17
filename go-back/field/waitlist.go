package field

import (
	"mobarter/app"
	"mobarter/database"

	"github.com/graphql-go/graphql"
)

func joinWaitList(app app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "waitListResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"email":     argString,
			"firstName": argString,
			"lastName":  argString,
			"country":   argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			result := app.DB.Create(&database.Waitlist{
				Email:     p.Args["email"].(string),
				FirstName: p.Args["firstName"].(string),
				LastName:  p.Args["lastName"].(string),
			})

			println("We put it all")
			if result.Error != nil {
				println(result.Error)

				return map[string]interface{}{
					"error": result.Error,
				}, nil
			}

			// return nil

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}

func JoinWaitList(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "waitListResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"email":     argString,
			"firstName": argString,
			"lastName":  argString,
			"country":   argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
