package handlers

import (
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
	"github.com/gin-gonic/gin"
)

// PoolHandler --> interface to Pool handler
type PoolHandler interface {
	GetPool(*gin.Context)
	GetAllPool(*gin.Context)
	AddPool(*gin.Context)
	UpdatePool(*gin.Context)
	DeletePool(*gin.Context)
}

type poolHandler struct {
	repo repository.PoolRepository
}

// NewPoolHandler --> returns new handler for Pool entity
func NewPoolHandler() PoolHandler {
	return &poolHandler{
		repo: repository.NewPoolRepository(),
	}
}

func (h *poolHandler) GetAllPool(ctx *gin.Context) {
	pool, err := h.repo.GetAllPool()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, pool)

}

func (h *poolHandler) GetPool(ctx *gin.Context) {
	poolStr := ctx.Param("Pool")
	poolID, err := strconv.Atoi(poolStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pool, err := h.repo.GetPool(poolID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, pool)

}

func (h *poolHandler) AddPool(ctx *gin.Context) {
	var pool models.Pool
	if err := ctx.ShouldBindJSON(&pool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pool, err := h.repo.AddPool(pool)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, pool)

}

func (h *poolHandler) UpdatePool(ctx *gin.Context) {

	var pool models.Pool
	if err := ctx.ShouldBindJSON(&pool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	poolStr := ctx.Param("Pool")
	poolID, err := strconv.Atoi(poolStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pool.ID = uint(poolID)
	pool, err = h.repo.UpdatePool(pool)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, pool)

}

func (h *poolHandler) DeletePool(ctx *gin.Context) {
	var pool models.Pool
	poolStr := ctx.Param("Pool")
	poolID, _ := strconv.Atoi(poolStr)
	pool.ID = uint(poolID)
	pool, err := h.repo.DeletePool(pool)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, pool)

}
