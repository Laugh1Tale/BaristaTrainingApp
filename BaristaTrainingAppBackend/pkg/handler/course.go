package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createCourse(c *gin.Context) {

}

// @Summary Courses
// @Tags api
// @Description courses list
// @ID courses
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/courses [get]
func (h *Handler) getCourses(c *gin.Context) {
	courses := []model.CourseResponse{
		{Id: 1, Name: "Стажировка", Description: "курс для того чтобы приступить к стажировке", Passing_Score: 8},
		{Id: 2, Name: "Бариста", Description: "курс для того чтобы приступить к работе бариста", Passing_Score: 9},
		{Id: 3, Name: "Менеджер", Description: "курс для того чтобы приступить к работе менеджера", Passing_Score: 10},
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getCourseById(c *gin.Context) {

}

func (h *Handler) updateCourse(c *gin.Context) {

}

func (h *Handler) deleteCourse(c *gin.Context) {

}
