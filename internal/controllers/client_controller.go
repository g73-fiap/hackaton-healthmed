package controllers

import (
	"errors"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientUseCase usecases.ClientUseCase
}

func (cc ClientController) GetClient(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	client, err := cc.clientUseCase.GetClient(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (cc ClientController) CreateClient(c *gin.Context) {
	var client entities.Client
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.clientUseCase.Createclient(client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (cc ClientController) UpdateClient(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	var client entities.Client
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.clientUseCase.UpdateClient(id, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (cc ClientController) DeleteClient(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	err := cc.clientUseCase.DeleteClient(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
