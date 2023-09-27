package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// ExtensionLogHandler --> interface to ExtensionLog handler
type ExtensionLogHandler interface {
	GetExtensionLog(*gin.Context)
	GetAllExtensionLog(*gin.Context)
	AddExtensionLog(*gin.Context)
	UpdateExtensionLog(*gin.Context)
	DeleteExtensionLog(*gin.Context)
}

type extensionLogHandler struct {
	repo repository.ExtensionLogRepository
}

// NewExtensionLogHandler --> returns new handler for ExtensionLog entity
func NewExtensionLogHandler() ExtensionLogHandler {
	return &extensionLogHandler{
		repo: repository.NewExtensionLogRepository(),
	}
}

func (h *extensionLogHandler) GetAllExtensionLog(ctx *gin.Context) {
	extensionLog, err := h.repo.GetAllExtensionLog()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, extensionLog)

}

func (h *extensionLogHandler) GetExtensionLog(ctx *gin.Context) {
	extensionLogStr := ctx.Param("extensionLog")
	extensionLogID, err := strconv.Atoi(extensionLogStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	extensionLog, err := h.repo.GetExtensionLog(extensionLogID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, extensionLog)

}

func (h *extensionLogHandler) AddExtensionLog(ctx *gin.Context) {
	var extensionLog models.ExtensionLog
	if err := ctx.ShouldBindJSON(&extensionLog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	extensionLog, err := h.repo.AddExtensionLog(extensionLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, extensionLog)

}

func (h *extensionLogHandler) UpdateExtensionLog(ctx *gin.Context) {

	var extensionLog models.ExtensionLog
	if err := ctx.ShouldBindJSON(&extensionLog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	extensionLogStr := ctx.Param("extensionLog")
	extensionLogID, err := strconv.Atoi(extensionLogStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	extensionLog.ID = uint(extensionLogID)
	extensionLog, err = h.repo.UpdateExtensionLog(extensionLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, extensionLog)

}

func (h *extensionLogHandler) DeleteExtensionLog(ctx *gin.Context) {
	var extensionLog models.ExtensionLog
	extensionLogStr := ctx.Param("extensionLog")
	extensionLogID, _ := strconv.Atoi(extensionLogStr)
	extensionLog.ID = uint(extensionLogID)
	extensionLog, err := h.repo.DeleteExtensionLog(extensionLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, extensionLog)

}
