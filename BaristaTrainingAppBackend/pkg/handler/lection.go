package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Lection struct {
	Id    int    `json:"-"`
	Theme string `json:"theme" binding:"required"`
}

func (h *Handler) getLections(c *gin.Context) {
	lections := []Lection{
		{Id: 1, Theme: "Латте"},
		{Id: 2, Theme: "Эспрессо"},
		{Id: 3, Theme: "Капучино"},
	}

	c.JSON(http.StatusOK, lections)
}
