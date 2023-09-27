package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// StationHandler --> interface to Station handler
type StationHandler interface {
	GetStation(*gin.Context)
	GetAllStation(*gin.Context)
	AddStation(*gin.Context)
	UpdateStation(*gin.Context)
	DeleteStation(*gin.Context)
}

type stationHandler struct {
	repo repository.StationRepository
}

// NewStationHandler --> returns new handler for Station entity
func NewStationHandler() StationHandler {
	return &stationHandler{
		repo: repository.NewStationRepository(),
	}
}

func (h *stationHandler) GetAllStation(ctx *gin.Context) {
	station, err := h.repo.GetAllStation()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, station)

}

func (h *stationHandler) GetStation(ctx *gin.Context) {
	stationStr := ctx.Param("Station")
	stationID, err := strconv.Atoi(stationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station, err := h.repo.GetStation(stationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, station)

}

func (h *stationHandler) AddStation(ctx *gin.Context) {
	var station models.Station
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station, err := h.repo.AddStation(station)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, station)

}

func (h *stationHandler) UpdateStation(ctx *gin.Context) {
	var station models.Station
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stationStr := ctx.Param("Station")
	stationID, err := strconv.Atoi(stationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station.ID = uint(stationID)
	station, err = h.repo.UpdateStation(station)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, station)

}

func (h *stationHandler) DeleteStation(ctx *gin.Context) {
	var station models.Station
	stationStr := ctx.Param("Station")
	stationID, _ := strconv.Atoi(stationStr)
	station.ID = uint(stationID)
	station, err := h.repo.DeleteStation(station)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, station)

}
