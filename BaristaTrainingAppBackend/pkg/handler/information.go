package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Information struct {
	Id    int    `json:"-"`
	Theme string `json:"theme" binding:"required"`
	Text  string `json:"Text" binding:"required"`
}

func (h *Handler) getInformations(c *gin.Context) {
	informations := []Information{
		{Id: 1, Theme: "Основы", Text: "Латте это кофе с молоком"},
		{Id: 2, Theme: "Пропорции", Text: "3 молока : 1 кофе"},
	}

	c.JSON(http.StatusOK, informations)
}
