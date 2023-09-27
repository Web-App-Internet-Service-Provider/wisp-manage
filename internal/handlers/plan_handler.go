package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PlanHandler --> interface to Plan handler
type PlanHandler interface {
	GetPlan(*gin.Context)
	GetAllPlan(*gin.Context)
	AddPlan(*gin.Context)
	UpdatePlan(*gin.Context)
	DeletePlan(*gin.Context)
}

type planHandler struct {
	repo repository.PlanRepository
}

// NewPlanHandler --> returns new handler for Plan entity
func NewPlanHandler() PlanHandler {
	return &planHandler{
		repo: repository.NewPlanRepository(),
	}
}

func (h *planHandler) GetAllPlan(ctx *gin.Context) {
	plan, err := h.repo.GetAllPlan()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, plan)

}

func (h *planHandler) GetPlan(ctx *gin.Context) {
	planStr := ctx.Param("plan")
	planID, err := strconv.Atoi(planStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := h.repo.GetPlan(planID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, plan)

}

func (h *planHandler) AddPlan(ctx *gin.Context) {
	var plan models.Plan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := h.repo.AddPlan(plan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, plan)

}

func (h *planHandler) UpdatePlan(ctx *gin.Context) {

	var plan models.Plan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planStr := ctx.Param("Plan")
	planID, err := strconv.Atoi(planStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan.ID = uint(planID)
	plan, err = h.repo.UpdatePlan(plan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, plan)

}

func (h *planHandler) DeletePlan(ctx *gin.Context) {
	var plan models.Plan
	planStr := ctx.Param("plan")
	planID, _ := strconv.Atoi(planStr)
	plan.ID = uint(planID)
	plan, err := h.repo.DeletePlan(plan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, plan)

}
