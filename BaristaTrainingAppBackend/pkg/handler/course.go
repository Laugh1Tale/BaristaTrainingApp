package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Id            int    `json:"-"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Passing_Score int    `json:"passing_score" binding:"required"`
}

func (h *Handler) getCourses(c *gin.Context) {
	courses := []Course{
		{Id: 1, Name: "Стажировка", Description: "курс для того чтобы приступить к стажировке", Passing_Score: 8},
		{Id: 2, Name: "Бариста", Description: "курс для того чтобы приступить к работе бариста", Passing_Score: 9},
		{Id: 3, Name: "Менеджер", Description: "курс для того чтобы приступить к работе менеджера", Passing_Score: 10},
	}

	c.JSON(http.StatusOK, courses)
}
