package main

import (
	"log"
	"net/http"
	"os"

	"blog-post-api/database"
	"blog-post-api/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	Database := database.Connect()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	customCtx := &graph.Resolver{
		DB: Database,
	}

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", graph.CreateContext(customCtx, srv))

	// create tables should be called only once, after that comment it
	database.CreateTables()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
