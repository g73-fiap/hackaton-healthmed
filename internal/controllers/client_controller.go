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

func NewClientController(clientUseCase usecases.ClientUseCase) ClientController {
	return ClientController{
		clientUseCase: clientUseCase,
	}
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

	err = cc.clientUseCase.UpdateClient(client)
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

func (cc ClientController) SaveMedicalReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.clientUseCase.SaveMedicalReport(id, file.Filename, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (cc ClientController) DeleteMedicalReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	fileName := c.Param("fileName")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("fileName is required")})
		return
	}

	err := cc.clientUseCase.RemoveMedicalReport(id, fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (cc ClientController) GetMedicalReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	fileName := c.Param("fileName")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("fileName is required")})
		return
	}

	var body struct {
		requester string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location, err := cc.clientUseCase.GetMedicalReport(id, fileName, body.requester)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fileLocation": location})
}

func (cc ClientController) ShareMedicalReport(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	fileName := c.Param("fileName")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("fileName is required")})
		return
	}

	var body struct {
		requester string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.clientUseCase.ShareMedicalReport(id, fileName, body.requester)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
