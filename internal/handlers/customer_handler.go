package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// CustomerHandler -> interface to Customer entity
type CustomerHandler interface {
	AddCustomer(*gin.Context)
	GetCustomer(*gin.Context)
	GetAllCustomer(*gin.Context)
	UpdateCustomer(*gin.Context)
	DeleteCustomer(*gin.Context)
	GetCustomerPlanType(*gin.Context)
}

type customerHandler struct {
	repo repository.CustomerRepository
}

// NewCustomerHandler --> returns new customer handler
func NewCustomerHandler() CustomerHandler {
	return &customerHandler{
		repo: repository.NewCustomerRepository(),
	}
}

func (h *customerHandler) GetAllCustomer(ctx *gin.Context) {
	fmt.Println(ctx.Get("customerID"))
	customer, err := h.repo.GetAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := h.repo.GetCustomer(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) AddCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := h.repo.AddCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	customer.Password = ""
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) UpdateCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("uscustomerer")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer.ID = uint(intID)
	customer, err = h.repo.UpdateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) DeleteCustomer(ctx *gin.Context) {
	var customer models.Customer
	id := ctx.Param("customer")
	intID, _ := strconv.Atoi(id)
	customer.ID = uint(intID)
	customer, err := h.repo.DeleteCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, customer)

}

func (h *customerHandler) GetCustomerPlanType(ctx *gin.Context) {
	customerStr := ctx.Param("customer")
	customerID, _ := strconv.Atoi(customerStr)
	if products, err := h.repo.GetCustomerPlanType(customerID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, products)
	}
}
