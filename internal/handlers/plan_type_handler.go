package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PlanTypeHandler --> interface to planType handler
type PlanTypeHandler interface {
	GetPlanType(*gin.Context)
	GetAllPlanType(*gin.Context)
	AddPlanType(*gin.Context)
	UpdatePlanType(*gin.Context)
	DeletePlanType(*gin.Context)
}

type planTypeHandler struct {
	repo repository.PlanTypeRepository
}

// NewplanTypeHandler --> returns new handler for planType entity
func NewPlanTypeHandler() PlanTypeHandler {
	return &planTypeHandler{
		repo: repository.NewPlanTypeRepository(),
	}
}

func (h *planTypeHandler) GetAllPlanType(ctx *gin.Context) {
	planType, err := h.repo.GetAllPlanType()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planType)

}

func (h *planTypeHandler) GetPlanType(ctx *gin.Context) {
	planTypeStr := ctx.Param("planType")
	planTypeID, err := strconv.Atoi(planTypeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planType, err := h.repo.GetPlanType(planTypeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planType)

}

func (h *planTypeHandler) AddPlanType(ctx *gin.Context) {
	var planType models.PlanType
	if err := ctx.ShouldBindJSON(&planType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planType, err := h.repo.AddPlanType(planType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planType)

}

func (h *planTypeHandler) UpdatePlanType(ctx *gin.Context) {

	var planType models.PlanType
	if err := ctx.ShouldBindJSON(&planType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planTypeStr := ctx.Param("planType")
	planTypeID, err := strconv.Atoi(planTypeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planType.ID = uint(planTypeID)
	planType, err = h.repo.UpdatePlanType(planType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planType)

}

func (h *planTypeHandler) DeletePlanType(ctx *gin.Context) {
	var planType models.PlanType
	planTypeStr := ctx.Param("planType")
	planTypeID, _ := strconv.Atoi(planTypeStr)
	planType.ID = uint(planTypeID)
	planType, err := h.repo.DeletePlanType(planType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planType)

}
