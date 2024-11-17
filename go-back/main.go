package main

import (
	"log"
	"net/http"

	"mobarter/app"
	"mobarter/database"
	"mobarter/field"
	"mobarter/helper"

	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/exp/slog"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		helper.ErrorPanic(err, "Could not load env")
	}

	config := &database.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)

	if err != nil {
		helper.ErrorPanic(err, "Cannot connect to db")
	}

	database.RunMigrations(db)
	// // Slog
	handlerOpt := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpt))

	slog.SetDefault(logger)

	appState := app.AppState{
		DB: db,
	}

	// fields := fieldsRegistry
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: field.QueryFields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutations", Fields: field.MutationsFields}

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

	// Start server
	http.Handle("/graphql", handler)
	log.Println("Server is running on http://localhost:" + os.Getenv("PORT") + "/graphql")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	server := appState.NewApp()

	appState.SetupRoutes(server, logger)

	// log.Fatal(server.Listen(":" + os.Getenv("PORT")))

}
