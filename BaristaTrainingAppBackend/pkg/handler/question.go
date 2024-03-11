package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Questions
// @Tags api
// @Description questions list
// @ID questions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/questions [get]
func (h *Handler) getQuestions(c *gin.Context) {
	questions := []model.QuestionResponse{
		{Id: 1, Theme: "Изготовление латте", Text: "С какими пропорциями эспрессо/молока варится латте", Right_answer_id: 1},
		{Id: 1, Theme: "Изготовление латте", Text: "Необходима ли молочная пена при изготовлении латте", Right_answer_id: 4},
		{Id: 1, Theme: "Сиропы для латте", Text: "Какого сиропа нет в нашей сети кофеен", Right_answer_id: 6},
	}

	c.JSON(http.StatusOK, questions)
}
