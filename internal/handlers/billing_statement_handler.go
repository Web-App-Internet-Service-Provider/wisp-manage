package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// BillingStatementHandler --> interface to billingStatement handler
type BillingStatementHandler interface {
	GetBillingStatement(*gin.Context)
	GetAllBillingStatement(*gin.Context)
	AddBillingStatement(*gin.Context)
	UpdateBillingStatement(*gin.Context)
	DeleteBillingStatement(*gin.Context)
}

type billingStatementHandler struct {
	repo repository.BillingStatementRepository
}

// NewBillingStatementHandler --> returns new handler for billingStatement entity
func NewBillingStatementHandler() BillingStatementHandler {
	return &billingStatementHandler{
		repo: repository.NewBillingStatementRepository(),
	}
}

func (h *billingStatementHandler) GetAllBillingStatement(ctx *gin.Context) {
	billingStatement, err := h.repo.GetAllBillingStatement()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, billingStatement)

}

func (h *billingStatementHandler) GetBillingStatement(ctx *gin.Context) {
	planTypeStr := ctx.Param("billingStatement")
	planTypeID, err := strconv.Atoi(planTypeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	billingStatement, err := h.repo.GetBillingStatement(planTypeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, billingStatement)

}

func (h *billingStatementHandler) AddBillingStatement(ctx *gin.Context) {
	var billingStatement models.BillingStatement
	if err := ctx.ShouldBindJSON(&billingStatement); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	billingStatement, err := h.repo.AddBillingStatement(billingStatement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, billingStatement)

}

func (h *billingStatementHandler) UpdateBillingStatement(ctx *gin.Context) {

	var billingStatement models.BillingStatement
	if err := ctx.ShouldBindJSON(&billingStatement); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planTypeStr := ctx.Param("billingStatement")
	planTypeID, err := strconv.Atoi(planTypeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	billingStatement.ID = uint(planTypeID)
	billingStatement, err = h.repo.UpdateBillingStatement(billingStatement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, billingStatement)

}

func (h *billingStatementHandler) DeleteBillingStatement(ctx *gin.Context) {
	var billingStatement models.BillingStatement
	planTypeStr := ctx.Param("billingStatement")
	planTypeID, _ := strconv.Atoi(planTypeStr)
	billingStatement.ID = uint(planTypeID)
	billingStatement, err := h.repo.DeleteBillingStatement(billingStatement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, billingStatement)

}
