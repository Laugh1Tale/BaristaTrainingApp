package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Id          int    `json:"-"`
	Theme       string `json:"theme" binding:"required"`
	Description string `json:"description" binding:"required"`
	Role_id     int    `json:"role_id" binding:"required"`
}

func (h *Handler) getTests(c *gin.Context) {
	tests := []Test{
		{Id: 1, Theme: "Латте", Description: "Тест по приготовлению латте", Role_id: 1},
	}

	c.JSON(http.StatusOK, tests)
}
