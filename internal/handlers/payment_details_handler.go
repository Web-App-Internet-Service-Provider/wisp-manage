package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PaymentDetailHandler --> interface to paymentDetail handler
type PaymentDetailHandler interface {
	GetPaymentDetail(*gin.Context)
	GetAllPaymentDetail(*gin.Context)
	AddPaymentDetail(*gin.Context)
	UpdatePaymentDetail(*gin.Context)
	DeletePaymentDetail(*gin.Context)
}

type paymentDetailHandler struct {
	repo repository.PaymentDetailRepository
}

// NewPaymentDetailHandler --> returns new handler for paymentDetail entity
func NewPaymentDetailHandler() PaymentDetailHandler {
	return &paymentDetailHandler{
		repo: repository.NewPaymentDetailRepository(),
	}
}

func (h *paymentDetailHandler) GetAllPaymentDetail(ctx *gin.Context) {
	paymentDetail, err := h.repo.GetAllPaymenDetail()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentDetail)

}

func (h *paymentDetailHandler) GetPaymentDetail(ctx *gin.Context) {
	paymentDetailStr := ctx.Param("paymentDetail")
	paymentDetailID, err := strconv.Atoi(paymentDetailStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentDetail, err := h.repo.GetPaymenDetail(paymentDetailID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentDetail)

}

func (h *paymentDetailHandler) AddPaymentDetail(ctx *gin.Context) {
	var paymentDetail models.PaymentDetail
	if err := ctx.ShouldBindJSON(&paymentDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentDetail, err := h.repo.AddPaymenDetail(paymentDetail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentDetail)

}

func (h *paymentDetailHandler) UpdatePaymentDetail(ctx *gin.Context) {

	var paymentDetail models.PaymentDetail
	if err := ctx.ShouldBindJSON(&paymentDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentDetailStr := ctx.Param("paymentDetail")
	paymentDetailID, err := strconv.Atoi(paymentDetailStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentDetail.ID = uint(paymentDetailID)
	paymentDetail, err = h.repo.UpdatePaymenDetail(paymentDetail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentDetail)

}

func (h *paymentDetailHandler) DeletePaymentDetail(ctx *gin.Context) {
	var paymentDetail models.PaymentDetail
	paymentDetailStr := ctx.Param("paymentDetail")
	paymentDetailID, _ := strconv.Atoi(paymentDetailStr)
	paymentDetail.ID = uint(paymentDetailID)
	paymentDetail, err := h.repo.DeletePaymenDetail(paymentDetail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, paymentDetail)

}
