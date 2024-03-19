package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Lections
// @Tags api
// @Description lections list
// @ID lections
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/lections [get]
func (h *Handler) getLections(c *gin.Context) {
	lections := []model.LectionResponse{
		{Id: 1, Theme: "Как варить латте"},
		{Id: 2, Theme: "Как варить эспрессо"},
		{Id: 3, Theme: "Как варить капучино"},
	}

	c.JSON(http.StatusOK, lections)
}
