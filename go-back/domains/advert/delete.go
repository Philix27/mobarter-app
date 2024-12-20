package user

import (
	"errors"
	"mobarter/app"
	"mobarter/db"
	"mobarter/domains/crypto"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/jackc/pgx/v5/pgtype"
)

func DeleteAd(ap app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Advert_DeleteAdsResponse",
			Fields: graphql.Fields{
				"message": app.String,
			},
		}),

		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "Advert_DeleteAdInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"token":  app.InputString,
							"adId":   app.InputString,
							"reason": app.InputStringOptional,
						},
					},
				),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			args, err := app.ValidateArg("input", p)
			if err != nil {
				return nil, errors.New("Input required")
			}
			// extract response data
			firstName := args["firstName"].(string)
			lastName := args["lastName"].(string)
			//TODO - update dob, add to db
			// dob := args["dob"].(string)
			token := args["token"].(string)

			err, payload := crypto.GetTokenPayload(token)
			if err != nil {
				return nil, errors.New("Invalid token")
			}

			i, err := strconv.Atoi(payload)
			if err != nil {
				return nil, errors.New("Invalid userId")
			}

			user, err := ap.DbQueries.User_GetById(ap.Ctx, int64(i))
			if err != nil {
				return nil, errors.New("User not found")
			}

			err = ap.DbQueries.User_Update(ap.Ctx, db.User_UpdateParams{
				ID: user.ID,
				FirstName: pgtype.Text{
					String: firstName,
				},
				LastName: pgtype.Text{
					String: lastName,
				},
			})

			if err != nil {
				println("Errox: " + err.Error())
				return nil, errors.New("Could not reset password")
			}

			return map[string]interface{}{
				"message": "success " + user.Email,
			}, nil

		},
	}
}
