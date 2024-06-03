package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson49/genproto/coffee"
	"log/slog"
	"net/http"
)

// BuyCoffee godoc
// @ID buy_coffee
// @Router /coffee [POST]
// @Summary Buy Coffee
// @Description Buy Coffee
// @Tags Coffee
// @Accept json
// @Produce json
// @Param coffee body coffee.BuyCoffee true "Coffee body"
// @Success 200 {object} coffee.PreparedCoffee "Coffee data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) BuyCoffee(ctx *gin.Context) {
	cf := coffee.BuyCoffee{}

	err := ctx.ShouldBindJSON(&cf)
	if err != nil {
		slog.Info("error coffee binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	coffee, err := h.srvs.CoffeeService.BuyingCoffee(ctx, &cf)
	if err != nil {
		slog.Info("error coffee binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
	}

	ctx.JSON(200, coffee)
}
