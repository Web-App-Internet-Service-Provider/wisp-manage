package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// OrganizationHandler -> interface to organization entity
type OrganizationHandler interface {
	AddOrganization(*gin.Context)
	GetOrganization(*gin.Context)
	GetAllOrganization(*gin.Context)
	UpdateOrganization(*gin.Context)
	DeleteOrganization(*gin.Context)
}

type organizationHandler struct {
	repo repository.OrganizationRepository
}

// NewOrganizationHandler --> returns new organization handler
func NewOrganizationHandler() OrganizationHandler {
	return &organizationHandler{
		repo: repository.NewOrganizationRepository(),
	}
}

func (h *organizationHandler) GetAllOrganization(ctx *gin.Context) {
	fmt.Println(ctx.Get("organizationID"))
	organization, err := h.repo.GetAllOrganization()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organization)

}

func (h *organizationHandler) GetOrganization(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organization, err := h.repo.GetOrganization(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organization)

}

func (h *organizationHandler) AddOrganization(ctx *gin.Context) {
	var organization models.Organization
	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organization, err := h.repo.AddOrganization(organization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, organization)

}

func (h *organizationHandler) UpdateOrganization(ctx *gin.Context) {
	var organization models.Organization
	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("organization")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	organization.ID = uint(intID)
	organization, err = h.repo.UpdateOrganization(organization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organization)

}

func (h *organizationHandler) DeleteOrganization(ctx *gin.Context) {
	var organization models.Organization
	id := ctx.Param("organization")
	intID, _ := strconv.Atoi(id)
	organization.ID = uint(intID)
	organization, err := h.repo.DeleteOrganization(organization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organization)

}
