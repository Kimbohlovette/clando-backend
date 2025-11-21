package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kimbohlovette/clando-backend/db/sqlc"
	"github.com/kimbohlovette/clando-backend/models"
)

func (s *server) createUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.store.Do().CreateUser(c, sqlc.CreateUserParams{
		ID:       uuid.New().String(),
		Username: req.Username,
		Phone:    req.Phone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (s *server) getUser(c *gin.Context) {
	user, err := s.store.Do().GetUser(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *server) getAllUsers(c *gin.Context) {
	users, err := s.store.Do().ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
