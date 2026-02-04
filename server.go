package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"test_crm/config"
	"test_crm/controllers"
	"test_crm/middleware"
	"test_crm/repository"
	"test_crm/routes"
	"test_crm/services"
	"test_crm/utils"
)

type Server struct {
	membershipRepo     repository.MembershipRepository
	contactRepo 	   repository.ContactRepository
	authService        services.AuthService
	membershipService  services.MembershipService
	contactService 	   services.ContactService
	jwtService         utils.JWTService
	db                 *sqlx.DB
	engine             *gin.Engine
	host               string
	cfg                *config.Config
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Username,
		cfg.DBConfig.Password,
		cfg.DBConfig.DBName,
	)

	db := sqlx.MustConnect(cfg.DBConfig.Driver, dsn)

	jwtService := utils.NewJWTService(cfg.JwtSignatureKey)
	membershipRepo := repository.NewMembershipRepository(db)
	contactRepo := repository.NewContactRepository(db)
	
	authService := services.NewAuthService(membershipRepo, jwtService)
	membershipService := services.NewMembershipService(membershipRepo)
	contactService := services.NewContactService(contactRepo)

	engine := gin.Default()

	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		membershipRepo:    membershipRepo,
		contactRepo: 	   contactRepo,
		authService:       authService,
		membershipService: membershipService,
		contactService:    contactService,
		jwtService:        jwtService,
		db:                db,
		engine:            engine,
		host:              host,
		cfg:               cfg,
	}
}

func (s *Server) SetupRoutes() {
	apiV1 := s.engine.Group("/api/v1")

	authMw := middleware.NewAuthMiddleware(s.jwtService)

	authController := controller.NewAuthController(s.authService)
	membershipController := controller.NewMembershipController(s.membershipService)
	contactController := controller.NewContactController(s.contactService)

	routes.AuthRoutes(apiV1, authController)
	routes.MembershipRoutes(apiV1, authMw, membershipController)
	routes.ContactRoutes(apiV1, authMw, contactController)
}

func (s *Server) Run() {
	s.SetupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatalf("failed to start server %s: %v", s.host, err)
	}
}

func (s *Server) Close() {
	if s.db != nil {
		s.db.Close()
	}
}