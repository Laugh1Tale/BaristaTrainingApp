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
		{Id: 0, Theme: "Латте"},
		{Id: 1, Theme: "Эспрессо"},
		{Id: 2, Theme: "Капучино"},
	}

	c.JSON(http.StatusOK, lections)
}
