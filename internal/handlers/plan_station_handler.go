package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PlanStationHandler --> interface to PlanStation handler
type PlanStationHandler interface {
	GetPlanStation(*gin.Context)
	GetAllPlanStation(*gin.Context)
	AddPlanStation(*gin.Context)
	UpdatePlanStation(*gin.Context)
	DeletePlanStation(*gin.Context)
}

type planStationHandler struct {
	repo repository.PlanStationRepository
}

// NewPlanStationHandler --> returns new handler for PlanStation entity
func NewPlanStationHandler() PlanStationHandler {
	return &planStationHandler{
		repo: repository.NewPlanStationRepository(),
	}
}

func (h *planStationHandler) GetAllPlanStation(ctx *gin.Context) {
	planStation, err := h.repo.GetAllPlanStation()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planStation)

}

func (h *planStationHandler) GetPlanStation(ctx *gin.Context) {
	planStationStr := ctx.Param("planStation")
	planStationID, err := strconv.Atoi(planStationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planStation, err := h.repo.GetPlanStation(planStationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planStation)

}

func (h *planStationHandler) AddPlanStation(ctx *gin.Context) {
	var planStation models.PlanStation
	if err := ctx.ShouldBindJSON(&planStation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planStation, err := h.repo.AddPlanStation(planStation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planStation)

}

func (h *planStationHandler) UpdatePlanStation(ctx *gin.Context) {

	var planStation models.PlanStation
	if err := ctx.ShouldBindJSON(&planStation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planStationStr := ctx.Param("planStation")
	planStationID, err := strconv.Atoi(planStationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planStation.ID = uint(planStationID)
	planStation, err = h.repo.UpdatePlanStation(planStation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planStation)

}

func (h *planStationHandler) DeletePlanStation(ctx *gin.Context) {
	var planStation models.PlanStation
	planStationStr := ctx.Param("PlanStation")
	planStationID, _ := strconv.Atoi(planStationStr)
	planStation.ID = uint(planStationID)
	planStation, err := h.repo.DeletePlanStation(planStation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, planStation)

}
