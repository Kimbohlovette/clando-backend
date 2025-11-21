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

func (s *server) createDriver(c *gin.Context) {
	var req models.Driver
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var rating pgtype.Numeric
	rating.Scan(req.Rating)
	
	driver, err := s.store.Do().CreateDriver(c, sqlc.CreateDriverParams{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Phone:       req.Phone,
		LicenseNo:   req.LicenseNo,
		VehicleType: req.VehicleType,
		VehicleNo:   req.VehicleNo,
		Rating:      rating,
		IsAvailable: &req.IsAvailable,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, driver)
}

func (s *server) getDriver(c *gin.Context) {
	driver, err := s.store.Do().GetDriver(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "driver not found"})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (s *server) getAllDrivers(c *gin.Context) {
	drivers, err := s.store.Do().ListAvailableDrivers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}
