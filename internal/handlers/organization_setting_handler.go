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
type OrganizationSettingHandler interface {
	AddOrganizationSetting(*gin.Context)
	GetOrganizationSetting(*gin.Context)
	GetAllOrganizationSetting(*gin.Context)
	UpdateOrganizationSetting(*gin.Context)
	DeleteOrganizationSetting(*gin.Context)
}

type organizationSettingHandler struct {
	repo repository.OrganizationSettingRepository
}

// NewOrganizationSettingHandler --> returns new OrganizationSetting handler
func NewOrganizationSettingHandler() OrganizationSettingHandler {
	return &organizationSettingHandler{
		repo: repository.NewOrganizationSettingRepository(),
	}
}

func (h *organizationSettingHandler) GetAllOrganizationSetting(ctx *gin.Context) {
	fmt.Println(ctx.Get("OrganizationSettingID"))
	organizationSetting, err := h.repo.GetAllOrganizationSetting()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organizationSetting)

}

func (h *organizationSettingHandler) GetOrganizationSetting(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizationSetting, err := h.repo.GetOrganizationSetting(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organizationSetting)

}

func (h *organizationSettingHandler) AddOrganizationSetting(ctx *gin.Context) {
	var organizationSetting models.OrganizationSetting
	if err := ctx.ShouldBindJSON(&organizationSetting); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizationSetting, err := h.repo.AddOrganizationSetting(organizationSetting)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, organizationSetting)

}

func (h *organizationSettingHandler) UpdateOrganizationSetting(ctx *gin.Context) {
	var organizationSetting models.OrganizationSetting
	if err := ctx.ShouldBindJSON(&organizationSetting); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("organizationSetting")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	organizationSetting.ID = uint(intID)
	organizationSetting, err = h.repo.UpdateOrganizationSetting(organizationSetting)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organizationSetting)

}

func (h *organizationSettingHandler) DeleteOrganizationSetting(ctx *gin.Context) {
	var organizationSetting models.OrganizationSetting
	id := ctx.Param("OrganizationSetting")
	intID, _ := strconv.Atoi(id)
	organizationSetting.ID = uint(intID)
	organizationSetting, err := h.repo.DeleteOrganizationSetting(organizationSetting)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, organizationSetting)

}
