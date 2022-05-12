package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/handler"
	"rxdrag.com/entify-schema-registry/authentication"
	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/schema"
)

const PORT = "8080"

func main() {
	config.Init()
	schema, err := schema.CreateSchema()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql",
		authentication.CorsMiddleware(
			authentication.AuthMiddleware(h),
		),
	)
	fmt.Println(fmt.Sprintf("Running a GraphQL API server at http://localhost:%s/graphql", PORT))
	err2 := http.ListenAndServe(":"+PORT, nil)
	if err2 != nil {
		fmt.Printf("Start failure:%s", err2)
	}
}
