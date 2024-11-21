package p2p

import (
	"mobarter/app"
	"mobarter/log"

	"github.com/graphql-go/graphql"
)

func GetDealer(appState app.AppState) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "GetDealerResponse",
			Fields: graphql.Fields{
				"name":           app.String,
				"id":             app.Int,
				"dexId":          app.Int,
				"upVotes":        app.Int,
				"downVotes":      app.Int,
				"totalVotes":     app.Int,
				"winPercentage":  app.Float,
				"lossPercentage": app.Float,
			},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ilog := log.New("Get bank account")
			ilog.Trace("Get bank account")
			return map[string]interface{}{
				"name": "Meat Pie",
			}, nil
		},
	}
}
