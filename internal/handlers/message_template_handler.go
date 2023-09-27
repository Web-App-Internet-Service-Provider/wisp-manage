package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// MessageTemplatesHandler --> interface to messageTemplate handler
type MessageTemplatesHandler interface {
	GetMessageTemplate(*gin.Context)
	GetAllMessageTemplate(*gin.Context)
	AddMessageTemplate(*gin.Context)
	UpdateMessageTemplate(*gin.Context)
	DeleteMessageTemplate(*gin.Context)
}

type messageTemplatesHandler struct {
	repo repository.MessageTemplateRepository
}

// NewMessageTemplatesHandler --> returns new handler for messageTemplate entity
func NewMessageTemplatesHandler() MessageTemplatesHandler {
	return &messageTemplatesHandler{
		repo: repository.NewMessagetemplateRepository(),
	}
}

func (h *messageTemplatesHandler) GetAllMessageTemplate(ctx *gin.Context) {
	messageTemplate, err := h.repo.GetAllMessageTemplate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageTemplate)

}

func (h *messageTemplatesHandler) GetMessageTemplate(ctx *gin.Context) {
	messageTemplateStr := ctx.Param("messageTemplate")
	messageTemplateID, err := strconv.Atoi(messageTemplateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageTemplate, err := h.repo.GetMessageTemplate(messageTemplateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageTemplate)

}

func (h *messageTemplatesHandler) AddMessageTemplate(ctx *gin.Context) {
	var messageTemplate models.MessageTemplate
	if err := ctx.ShouldBindJSON(&messageTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageTemplate, err := h.repo.AddMessageTemplate(messageTemplate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageTemplate)

}

func (h *messageTemplatesHandler) UpdateMessageTemplate(ctx *gin.Context) {

	var messageTemplate models.MessageTemplate
	if err := ctx.ShouldBindJSON(&messageTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageTemplateStr := ctx.Param("messageTemplate")
	messageTemplateID, err := strconv.Atoi(messageTemplateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageTemplate.ID = uint(messageTemplateID)
	messageTemplate, err = h.repo.UpdateMessageTemplate(messageTemplate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageTemplate)

}

func (h *messageTemplatesHandler) DeleteMessageTemplate(ctx *gin.Context) {
	var messageTemplate models.MessageTemplate
	messageTemplateStr := ctx.Param("messageTemplate")
	messageTemplateID, _ := strconv.Atoi(messageTemplateStr)
	messageTemplate.ID = uint(messageTemplateID)
	messageTemplate, err := h.repo.DeleteMessageTemplate(messageTemplate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, messageTemplate)

}
