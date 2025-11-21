package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kimbohlovette/clando-backend/db"
)

type Server interface {
	Start(address string) error
}

type server struct {
	store  db.AppStore
	router *gin.Engine
}

func NewServer(store db.AppStore) Server {
	s := &server{
		store:  store,
		router: gin.Default(),
	}
	s.setupRoutes()
	return s
}

func (s *server) setupRoutes() {
	api := s.router.Group("/api")
	
	api.POST("/users", s.createUser)
	api.GET("/users/:id", s.getUser)
	api.GET("/users", s.getAllUsers)
	
	api.POST("/payments/initiate", s.initiatePayment)
	api.PUT("/payments/:id", s.updatePayment)
	api.PUT("/payments/:id/status", s.updatePaymentStatus)
	api.GET("/payments", s.getAllPayments)
	
	api.POST("/places", s.createPlace)
	api.GET("/places/:id", s.getPlace)
	api.GET("/places", s.getAllPlaces)
	
	api.POST("/drivers", s.createDriver)
	api.GET("/drivers/:id", s.getDriver)
	api.GET("/drivers", s.getAllDrivers)
	
	api.POST("/calculate-fare", s.calculateFare)
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
