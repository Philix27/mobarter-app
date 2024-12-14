package main

import (
	"log"
	"net/http"

	"mobarter/app"
	"mobarter/config"

	"mobarter/db"
	field "mobarter/domains"

	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
	"golang.org/x/exp/slog"
)

func main() {

	config := config.SetupEnv()

	handlerOpt := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpt))

	slog.SetDefault(logger)

	queries, ctx, conn := db.Connect(config.DbUrl)
	
	defer conn.Close(ctx)

	appState := app.AppState{
		Logger:    logger,
		DbQueries: queries,
		Ctx:       ctx,
	}

	// fields := fieldsRegistry
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: field.QueryFields(appState)}
	rootMutation := graphql.ObjectConfig{Name: "RootMutations", Fields: field.MutationsFields(appState)}

	// Schema
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// handler
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		Debug:          true,
	})

	handler := c.Handler(h)

	port := ":" + os.Getenv("PORT")
	// Start server
	http.Handle("/graphql", handler)
	log.Println("Server is running on http://localhost" + port + "/graphql")
	log.Fatal(http.ListenAndServe(port, nil))

}
