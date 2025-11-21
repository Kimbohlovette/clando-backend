package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kimbohlovette/clando-backend/db/sqlc"
	"github.com/kimbohlovette/clando-backend/models"
)

func (s *server) createPlace(c *gin.Context) {
	var req models.Place
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var lat, lon pgtype.Numeric
	lat.Scan(req.Latitude)
	lon.Scan(req.Longitude)
	
	place, err := s.store.Do().CreatePlace(c, sqlc.CreatePlaceParams{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Address:   req.Address,
		Latitude:  lat,
		Longitude: lon,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, place)
}

func (s *server) getPlace(c *gin.Context) {
	place, err := s.store.Do().GetPlace(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "place not found"})
		return
	}
	c.JSON(http.StatusOK, place)
}

func (s *server) getAllPlaces(c *gin.Context) {
	places, err := s.store.Do().ListPlaces(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, places)
}
