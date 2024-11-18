package field

import (
	"fmt"
	"mobarter/app"
	"mobarter/database"

	"github.com/graphql-go/graphql"
)

func JoinWaitList(app app.AppState) *graphql.Field {
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
			fmt.Println("Hit waitlist list")
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
				}, result.Error
			}

			// return nil

			return map[string]interface{}{
				"message": "success",
			}, nil
		},
	}
}
