package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// LocationHandler -> interface to location entity
type LocationHandler interface {
	AddLocation(*gin.Context)
	GetLocation(*gin.Context)
	GetAllLocation(*gin.Context)
	UpdateLocation(*gin.Context)
	DeleteLocation(*gin.Context)
}

type locationHandler struct {
	repo repository.LocationRepository
}

// NewLocationHandler --> returns new location handler
func NewLocationHandler() LocationHandler {
	return &locationHandler{
		repo: repository.NewLocationRepository(),
	}
}

func (h *locationHandler) GetAllLocation(ctx *gin.Context) {
	fmt.Println(ctx.Get("locationID"))
	location, err := h.repo.GetAllLocation()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, location)

}

func (h *locationHandler) GetLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location, err := h.repo.GetLocation(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, location)

}

func (h *locationHandler) AddLocation(ctx *gin.Context) {
	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location, err := h.repo.AddLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, location)

}

func (h *locationHandler) UpdateLocation(ctx *gin.Context) {
	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("location")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	location.ID = uint(intID)
	location, err = h.repo.UpdateLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, location)

}

func (h *locationHandler) DeleteLocation(ctx *gin.Context) {
	var location models.Location
	id := ctx.Param("location")
	intID, _ := strconv.Atoi(id)
	location.ID = uint(intID)
	location, err := h.repo.DeleteLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, location)

}
