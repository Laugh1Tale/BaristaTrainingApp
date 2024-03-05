package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Lesson struct {
// 	Id            int    `json:"-"`
// 	Name          string `json:"name" binding:"required"`
// 	Description   string `json:"description" binding:"required"`
// 	Passing_Score int    `json:"passing_score" binding:"required"`
// }

// func (h *Handler) getLessons(c *gin.Context) {
// 	lessons := []Lesson{
// 		{Id: 1, Name: "Как варить кофе", Description: "Изготовление разных видов кофе", Passing_Score: 8},
// 		{Id: 2, Name: "Клиентоориентированность", Description: "Как общаться c клиентами", Passing_Score: 8},
// 		{Id: 3, Name: "Стрессоустойчивость", Description: "Умение справляться в стрессовых ситуациях", Passing_Score: 8},
// 	}

// 	c.JSON(http.StatusOK, lessons)
// ?
