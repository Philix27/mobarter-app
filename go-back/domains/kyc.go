package field

import (
	"mobarter/app"

	"github.com/graphql-go/graphql"
)

func Kyc_VerifyNin(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_VerifyNinResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"nin": argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
func Kyc_VerifyBvn(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_VerifyBvnResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"bvn": argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
func Kyc_VerifyEmail(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_VerifyEmailResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"email": argString,
			"otp":   argString,
			"token": argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
func Kyc_VerifyPhone(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_VerifyPhoneResponse",
			Fields: graphql.Fields{
				"message": String,
			},
		}),
		Args: graphql.FieldConfigArgument{
			"phone": argString,
			"otp":   argString,
			"token": argString,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}

func Kyc_GetLevel(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Kyc_GetLevelResponse",
			Fields: graphql.Fields{
				"accountName": String,
				"accountNo":   String,
				"bankName":    String,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
