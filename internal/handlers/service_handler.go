package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// ServiceHandler --> interface to Service handler
type ServiceHandler interface {
	GetService(*gin.Context)
	GetAllService(*gin.Context)
	AddService(*gin.Context)
	UpdateService(*gin.Context)
	DeleteService(*gin.Context)
}

type serviceHandler struct {
	repo repository.ServiceRepository
}

// NewServiceHandler --> returns new handler for Service entity
func NewServiceHandler() ServiceHandler {
	return &serviceHandler{
		repo: repository.NewServiceRepository(),
	}
}

func (h *serviceHandler) GetAllService(ctx *gin.Context) {
	service, err := h.repo.GetAllService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, service)

}

func (h *serviceHandler) GetService(ctx *gin.Context) {
	serviceStr := ctx.Param("Service")
	serviceID, err := strconv.Atoi(serviceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service, err := h.repo.GetService(serviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, service)

}

func (h *serviceHandler) AddService(ctx *gin.Context) {
	var service models.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service, err := h.repo.AddService(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, service)

}

func (h *serviceHandler) UpdateService(ctx *gin.Context) {

	var service models.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serviceStr := ctx.Param("service")
	serviceID, err := strconv.Atoi(serviceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.ID = uint(serviceID)
	service, err = h.repo.UpdateService(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, service)

}

func (h *serviceHandler) DeleteService(ctx *gin.Context) {
	var service models.Service
	serviceStr := ctx.Param("Service")
	serviceID, _ := strconv.Atoi(serviceStr)
	service.ID = uint(serviceID)
	service, err := h.repo.DeleteService(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, service)

}
