package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Answer struct {
	Id   int    `json:"-"`
	Text string `json:"text" binding:"required"`
}

func (h *Handler) getAnswers(c *gin.Context) {
	answers := []Answer{
		{Id: 1, Text: "3 молока : 1 эспрессо"},
		{Id: 2, Text: "1 молоко : 3 эспрессо"},
		{Id: 3, Text: "нет (не нужно)"},
		{Id: 4, Text: "да (необходимо)"},
		{Id: 5, Text: "миндальный"},
		{Id: 6, Text: "абрикосовый"},
	}

	c.JSON(http.StatusOK, answers)
}
