// main.go
package main

import (
	"adr/backend/src/graphql/generated"
	graph "adr/backend/src/graphql/resolver"
	"adr/backend/src/middleware"
	"adr/backend/src/model"
	"adr/backend/src/utils"
	"context"
	"flag"
	"log"
	"os"

	"github.com/arsmn/fastgql/graphql"
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var app *fiber.App

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")

	godotenv.Load(".env")
	app = fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api", middleware.Auth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			"Everything is up and running",
		)
	})

	api.All("/graphql", func(c *fiber.Ctx) error {
		config := generated.Config{
			Resolvers: &graph.Resolver{},
		}

		client, clientCtx, err := utils.PrismaClient()
		if err != nil {
			return err
		}

		defer func() {
			if err := client.Prisma.Disconnect(); err != nil {
				log.Println(err)
			}
		}()

		user := c.Locals("user")
		accessToken := c.Locals("accessToken")
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
		srv.AroundFields(
			func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				ctx = context.WithValue(ctx, model.ConfigKey("user"), user)
				ctx = context.WithValue(ctx, model.ConfigKey("accessToken"), accessToken)
				ctx = context.WithValue(ctx, model.ConfigKey("client"), client)
				ctx = context.WithValue(ctx, model.ConfigKey("clientCtx"), clientCtx)
				return next(ctx)
			},
		)
		gqlHandler := srv.Handler()
		gqlHandler(c.Context())
		return nil
	})

}

var (
	addr = flag.String("addr :", os.Getenv("PORT"), "")
)

func main() {

	sentryUrl := os.Getenv("SENTRY")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryUrl,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	flag.Parse()

	if *addr == "" {
		*addr = ":9000"
	}
	err = app.Listen(*addr)

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Fiber cold start")
	}
}
