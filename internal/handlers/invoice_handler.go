package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// InvoiceHandler -> interface to invoice entity
type InvoiceHandler interface {
	AddInvoice(*gin.Context)
	GetInvoice(*gin.Context)
	GetAllInvoice(*gin.Context)
	UpdateInvoice(*gin.Context)
	DeleteInvoice(*gin.Context)
}

type invoiceHandler struct {
	repo repository.InvoiceRepository
}

// NewInvoiceHandler --> returns new invoice handler
func NewInvoiceHandler() InvoiceHandler {
	return &invoiceHandler{
		repo: repository.NewInvoiceRepository(),
	}
}

func (h *invoiceHandler) GetAllInvoice(ctx *gin.Context) {
	fmt.Println(ctx.Get("invoiceID"))
	invoice, err := h.repo.GetAllInvoice()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, invoice)

}

func (h *invoiceHandler) GetInvoice(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoice, err := h.repo.GetInvoice(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, invoice)

}

func (h *invoiceHandler) AddInvoice(ctx *gin.Context) {
	var invoice models.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoice, err := h.repo.AddInvoice(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, invoice)

}

func (h *invoiceHandler) UpdateInvoice(ctx *gin.Context) {
	var invoice models.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("invoice")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	invoice.ID = uint(intID)
	invoice, err = h.repo.UpdateInvoice(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, invoice)

}

func (h *invoiceHandler) DeleteInvoice(ctx *gin.Context) {
	var invoice models.Invoice
	id := ctx.Param("invoice")
	intID, _ := strconv.Atoi(id)
	invoice.ID = uint(intID)
	invoice, err := h.repo.DeleteInvoice(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, invoice)

}
