package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Tests
// @Tags api
// @Description tests list
// @ID tests
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/tests [get]
func (h *Handler) getTests(c *gin.Context) {
	tests := []model.TestResponse{
		{Id: 1, Theme: "Латте", Description: "Тест по приготовлению латте", Role_id: 1},
	}

	c.JSON(http.StatusOK, tests)
}
