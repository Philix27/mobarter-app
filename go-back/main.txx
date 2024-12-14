package main

import (
	"log"
	"net/http"

	"mobarter/app"
	"mobarter/database"
	field "mobarter/domains"

	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/exp/slog"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		app.ErrorPanic(err, "Could not load env")
	}
db.Queries
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
		app.ErrorPanic(err, "Cannot connect to db")
	}

	// database.RunMigrations(db)
	// // Slog
	handlerOpt := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpt))

	slog.SetDefault(logger)

	appState := app.AppState{
		DB:     db,
		Logger: logger,
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
