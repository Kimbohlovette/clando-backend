package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kimbohlovette/clando-backend/db/sqlc"
	"github.com/kimbohlovette/clando-backend/models"
)

func (s *server) initiatePayment(c *gin.Context) {
	var req models.Payment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var amount pgtype.Numeric
	amount.Scan(req.Amount)
	
	payment, err := s.store.Do().CreatePayment(c, sqlc.CreatePaymentParams{
		ID:            req.ID,
		UserID:        req.UserID,
		TravelID:      req.TravelID,
		Amount:        amount,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
		CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, payment)
}

func (s *server) updatePayment(c *gin.Context) {
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := s.store.Do().UpdatePaymentStatus(c, sqlc.UpdatePaymentStatusParams{
		ID:     c.Param("id"),
		Status: req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "payment updated"})
}

func (s *server) updatePaymentStatus(c *gin.Context) {
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := s.store.Do().UpdatePaymentStatus(c, sqlc.UpdatePaymentStatusParams{
		ID:     c.Param("id"),
		Status: req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "payment status updated"})
}

func (s *server) getAllPayments(c *gin.Context) {
	userID := c.Query("user_id")
	if userID != "" {
		payments, err := s.store.Do().ListUserPayments(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, payments)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "user_id query parameter required"})
}
