package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// MessageLogsHandler --> interface to messageLog handler
type MessageLogsHandler interface {
	GetMessageLog(*gin.Context)
	GetAllMessageLog(*gin.Context)
	AddMessageLog(*gin.Context)
	UpdateMessageLog(*gin.Context)
	DeleteMessageLog(*gin.Context)
}

type messageLogsHandler struct {
	repo repository.MessagelogRepository
}

// NewMessageLogsHandler --> returns new handler for messageLog entity
func NewMessageLogsHandler() MessageLogsHandler {
	return &messageLogsHandler{
		repo: repository.NewMessagelogRepository(),
	}
}

func (h *messageLogsHandler) GetAllMessageLog(ctx *gin.Context) {
	messageLog, err := h.repo.GetAllMessageLog()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageLog)

}

func (h *messageLogsHandler) GetMessageLog(ctx *gin.Context) {
	messageLogStr := ctx.Param("messageLog")
	messageLogID, err := strconv.Atoi(messageLogStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageLog, err := h.repo.GetMessageLog(messageLogID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageLog)

}

func (h *messageLogsHandler) AddMessageLog(ctx *gin.Context) {
	var messageLog models.MessageLog
	if err := ctx.ShouldBindJSON(&messageLog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageLog, err := h.repo.AddMessageLog(messageLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageLog)

}

func (h *messageLogsHandler) UpdateMessageLog(ctx *gin.Context) {

	var messageLog models.MessageLog
	if err := ctx.ShouldBindJSON(&messageLog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageLogStr := ctx.Param("messageLog")
	messageLogID, err := strconv.Atoi(messageLogStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageLog.ID = uint(messageLogID)
	messageLog, err = h.repo.UpdateMessageLog(messageLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageLog)

}

func (h *messageLogsHandler) DeleteMessageLog(ctx *gin.Context) {
	var messageLog models.MessageLog
	messageLogStr := ctx.Param("messageLog")
	messageLogID, _ := strconv.Atoi(messageLogStr)
	messageLog.ID = uint(messageLogID)
	messageLog, err := h.repo.DeleteMessageLog(messageLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageLog)

}
