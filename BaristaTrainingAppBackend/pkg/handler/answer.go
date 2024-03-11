package handler

import (
	"barTrApp/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Answers
// @Tags api
// @Description Answers list
// @ID answers
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Router /api/answers [get]
func (h *Handler) getAnswers(c *gin.Context) {
	answers := []model.AnswerResponse{
		{Id: 1, Text: "3 молока : 1 эспрессо"},
		{Id: 2, Text: "1 молоко : 3 эспрессо"},
		{Id: 3, Text: "нет (не нужно)"},
		{Id: 4, Text: "да (необходимо)"},
		{Id: 5, Text: "миндальный"},
		{Id: 6, Text: "абрикосовый"},
	}

	c.JSON(http.StatusOK, answers)
}
