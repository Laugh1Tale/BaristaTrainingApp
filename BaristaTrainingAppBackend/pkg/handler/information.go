package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Information struct {
	Id    int    `json:"-"`
	Theme string `json:"theme" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

func (h *Handler) getInformations(c *gin.Context) {
	informations := []Information{
		{Id: 1, Theme: "Основы", Text: "Латте это кофе с молоком и молочной пеной"},
		{Id: 2, Theme: "Пропорции", Text: "3 молока : 1 кофе"},
		{Id: 3, Theme: "Сиропы", Text: "В нашей сети кофеен есть только миндальный сироп"},
	}

	c.JSON(http.StatusOK, informations)
}
