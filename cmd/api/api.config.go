package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pinjoc/pinjoc-backend/internal/handler"
	"github.com/pinjoc/pinjoc-backend/lib/config"
)

type Application struct {
	app    *fiber.App
	appCfg AppConfig
}

type AppConfig struct {
	h handler.Handler
	c config.Config
}

func NewApp(config AppConfig) *Application {
	return &Application{
		app:    fiber.New(),
		appCfg: config,
	}
}

func (a *Application) RegisterRoute() {
	a.app.Use(recover.New())
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
	}))

	apiRouter := a.app.Group("api/v1")

	clobRouter := apiRouter.Group("/clob")
	clobRouter.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	clobRouter.Get("/clob", a.appCfg.h.CLOB.GetCLOB)
	clobRouter.Get("/available-token", a.appCfg.h.CLOB.GetAvailableToken)
	clobRouter.Get("/best-rate", a.appCfg.h.CLOB.GetBestRate)

	tokenRouter := apiRouter.Group("/token")
	tokenRouter.Get("/token", a.appCfg.h.Tokenized.GetToken)
	tokenRouter.Get("/available-token", a.appCfg.h.Tokenized.GetAllToken)
	tokenRouter.Get("/best-price", a.appCfg.h.Tokenized.GetBestPrice)
}

func (a *Application) Run() error {
	log.Println("server running")
	return a.app.Listen(a.appCfg.c.AddrHttp)
}
