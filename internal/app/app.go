package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	db "github.com/kyusupov33/mm-go-service/database"
	"github.com/valyala/fasthttp"
)

type Application struct {
	ctx context.Context
}

func New() *Application {
	ctx := context.Background()

	initHTTPHandlers(ctx)

	app := &Application{
		ctx: ctx,
	}
	return app
}

func (a *Application) Start() {
	dbName := os.Getenv("ENV.DB_NAME")
	dbPass := os.Getenv("ENV.DB_PASS")
	dbHost := os.Getenv("ENV.DB_HOST")
	dbPort := 5432

	connection, error := db.ConnectSQL(dbHost, dbPort, dbName, dbPass, dbName)
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func initHTTPHandlers(ctx context.Context) {
	// apiService := api.New(ctx)

	// authMw := utils.AuthMiddleware(authorizer)
}

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}
