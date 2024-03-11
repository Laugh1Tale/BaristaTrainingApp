package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Information
// @Tags api
// @Description information list
// @ID information
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/information [get]
func (h *Handler) getInformations(c *gin.Context) {
	informations := []model.InformationResponse{
		{Id: 1, Theme: "Основы", Text: "Латте это кофе с молоком и молочной пеной"},
		{Id: 2, Theme: "Пропорции", Text: "3 молока : 1 кофе"},
		{Id: 3, Theme: "Сиропы", Text: "В нашей сети кофеен есть только миндальный сироп"},
	}

	c.JSON(http.StatusOK, informations)
}
