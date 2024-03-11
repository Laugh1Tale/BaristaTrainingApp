package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Lessons
// @Tags api
// @Description Lessons list
// @ID lessons
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/lessons [get]
func (h *Handler) getLessons(c *gin.Context) {
	lessons := []model.LessonResponse{
		{Id: 1, Name: "Как варить кофе", Description: "Изготовление разных видов кофе", Passing_Score: 8},
		{Id: 2, Name: "Клиентоориентированность", Description: "Как общаться c клиентами", Passing_Score: 8},
		{Id: 3, Name: "Стрессоустойчивость", Description: "Умение справляться в стрессовых ситуациях", Passing_Score: 8},
	}

	c.JSON(http.StatusOK, lessons)
}
