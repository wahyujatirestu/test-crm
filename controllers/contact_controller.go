package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"test_crm/dto"
	"test_crm/services"
)

type ContactController struct {
	service services.ContactService
}

func NewContactController(s services.ContactService) *ContactController {
	return &ContactController{s}
}

func (c *ContactController) Create(ctx *gin.Context) {
	membershipID, _ := strconv.Atoi(ctx.Param("id"))

	var req dto.CreateContactRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(ctx, membershipID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "contact created"})
}

func (c *ContactController) GetByMembership(ctx *gin.Context) {
	membershipID, _ := strconv.Atoi(ctx.Param("membershipId"))

	data, err := c.service.GetByMembership(ctx, membershipID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *ContactController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req dto.UpdateContactRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Update(ctx, id, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact updated"})
}

func (c *ContactController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact deleted"})
}
