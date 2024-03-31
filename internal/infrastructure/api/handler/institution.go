package handler

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type InstitutionHandler struct {
	is services.InstitutionService
}

type institutionBody struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Country  string `json:"country"  binding:"required"`
	Province string `json:"province" binding:"required"`
	City     string `json:"city"     binding:"required"`
	Code     string `json:"code"     binding:"required"`
	Language string `json:"language" binding:"required"`
}

type institutionParams struct {
	ID uint64 `uri:"id" binding:"required"`
}

func NewInstitutionHandler(is services.InstitutionService) *InstitutionHandler {
	return &InstitutionHandler{is: is}
}

func (ih *InstitutionHandler) CreateInstitution(c *gin.Context) {
	var institution institutionBody
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&institution); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	if err := ih.is.CreateInstitution(ctx, toDomainInstitution(institution)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(201, gin.H{"created": true})
}

func (ih *InstitutionHandler) GetInstitutions(c *gin.Context) {
	ctx := c.Request.Context()
	institutions, err := ih.is.GetInstitutions(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{"institutions": institutions})
}

func (ih *InstitutionHandler) DeleteInstitution(c *gin.Context) {
	var params institutionParams
	ctx := c.Request.Context()

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	if err := ih.is.DeleteInstitution(ctx, params.ID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{"deleted": true})
}

func toDomainInstitution(institution institutionBody) domain.Institution {
	return domain.Institution{
		Name:     institution.Name,
		Email:    institution.Email,
		Country:  institution.Country,
		Province: institution.Province,
		City:     institution.City,
		Code:     institution.Code,
		Language: institution.Language,
	}
}
