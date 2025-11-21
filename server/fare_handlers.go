package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) calculateFare(c *gin.Context) {
	var req struct {
		Distance    float64 `json:"distance"`
		VehicleType string  `json:"vehicle_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	baseFare := 500.0
	perKmRate := 200.0
	
	if req.VehicleType == "premium" {
		perKmRate = 300.0
	}
	
	fare := baseFare + (req.Distance * perKmRate)
	
	c.JSON(http.StatusOK, gin.H{
		"distance":     req.Distance,
		"vehicle_type": req.VehicleType,
		"fare":         fare,
	})
}
