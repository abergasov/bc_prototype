package routes

import (
	"bc_prototype/internal/logger"
	"bc_prototype/internal/service/circle"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

var testCircle = uuid.MustParse("372e3c13-a552-4ebe-96fb-7f1bf45cf68c")

type Server struct {
	appAddr    string
	log        logger.AppLogger
	service    *circle.Service
	httpEngine *fiber.App
}

// InitAppRouter initializes the HTTP Server.
func InitAppRouter(log logger.AppLogger, service *circle.Service, address string) *Server {
	app := &Server{
		appAddr:    address,
		httpEngine: fiber.New(fiber.Config{}),
		service:    service,
		log:        log.With(zap.String("service", "http")),
	}
	app.httpEngine.Use(recover.New())
	app.initRoutes()
	return app
}

func (s *Server) initRoutes() {
	s.httpEngine.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	apiGroup := s.httpEngine.Group("/api")
	{
		apiGroup.Get("/members", s.listOfMembers)
		apiGroup.Post("/members", s.saveMemberAddress)
		apiGroup.Get("/circle/:circleID/contract", s.getCircleContract)
		apiGroup.Post("/circle/:circleID/contract", s.deployContract)
	}
}

// Run starts the HTTP Server.
func (s *Server) Run() error {
	s.log.Info("Starting HTTP server", zap.String("port", s.appAddr))
	return s.httpEngine.Listen(s.appAddr)
}

func (s *Server) Stop() error {
	return s.httpEngine.Shutdown()
}
