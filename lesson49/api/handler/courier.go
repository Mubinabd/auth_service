package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson49/genproto/courier"
	"log/slog"
	"net/http"
)

// Delivering godoc
// @ID deliver
// @Router /courier [POST]
// @Summary Buy Deliver
// @Description Buy Deliver
// @Tags Deliver
// @Accept json
// @Produce json
// @Param user body courier.TakeOrder true "Coffee body"
// @Success 200 {object} courier.DeliverOrder "Coffee data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Delivering(ctx *gin.Context) {
	cf := courier.TakeOrder{}

	err := ctx.ShouldBindJSON(&cf)
	if err != nil {
		slog.Info("error courier binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	courier, err := h.srvs.CourierService.Deliver(ctx, &cf)
	if err != nil {
		slog.Info("error courier binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
	}

	ctx.JSON(200, courier)
}
