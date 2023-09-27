package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PaymentMethodHandler --> interface to PaymentMethod handler
type PaymentMethodHandler interface {
	GetPaymentMethod(*gin.Context)
	GetAllPaymentMethod(*gin.Context)
	AddPaymentMethod(*gin.Context)
	UpdatePaymentMethod(*gin.Context)
	DeletePaymentMethod(*gin.Context)
}

type paymentMethodHandler struct {
	repo repository.PaymentMethodRepository
}

// NewPaymentMethodHandler --> returns new handler for PaymentMethod entity
func NewPaymentMethodHandler() PaymentMethodHandler {
	return &paymentMethodHandler{
		repo: repository.NewPaymentMethodRepository(),
	}
}

func (h *paymentMethodHandler) GetAllPaymentMethod(ctx *gin.Context) {
	paymentMethod, err := h.repo.GetAllPaymenMethod()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentMethod)

}

func (h *paymentMethodHandler) GetPaymentMethod(ctx *gin.Context) {
	paymentMethodStr := ctx.Param("PaymentMethod")
	paymentMethodID, err := strconv.Atoi(paymentMethodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentMethod, err := h.repo.GetPaymenMethod(paymentMethodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentMethod)

}

func (h *paymentMethodHandler) AddPaymentMethod(ctx *gin.Context) {
	var paymentMethod models.PaymentMethod
	if err := ctx.ShouldBindJSON(&paymentMethod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentMethod, err := h.repo.AddPaymenMethod(paymentMethod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentMethod)

}

func (h *paymentMethodHandler) UpdatePaymentMethod(ctx *gin.Context) {

	var paymentMethod models.PaymentMethod
	if err := ctx.ShouldBindJSON(&paymentMethod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentMethodStr := ctx.Param("PaymentMethod")
	paymentMethodID, err := strconv.Atoi(paymentMethodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentMethod.ID = uint(paymentMethodID)
	paymentMethod, err = h.repo.UpdatePaymenMethod(paymentMethod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentMethod)

}

func (h *paymentMethodHandler) DeletePaymentMethod(ctx *gin.Context) {
	var paymentMethod models.PaymentMethod
	paymentMethodStr := ctx.Param("PaymentMethod")
	paymentMethodID, _ := strconv.Atoi(paymentMethodStr)
	paymentMethod.ID = uint(paymentMethodID)
	paymentMethod, err := h.repo.DeletePaymenMethod(paymentMethod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentMethod)

}
