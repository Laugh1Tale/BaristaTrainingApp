package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Question struct {
	Id              int    `json:"-"`
	Theme           string `json:"theme" binding:"required"`
	Text            string `json:"text" binding:"required"`
	Right_answer_id int    `json:"right_answer_id" binding:"required"`
}

func (h *Handler) getQuestions(c *gin.Context) {
	questions := []Question{
		{Id: 1, Theme: "Изготовление латте", Text: "С какими пропорциями эспрессо/молока варится латте", Right_answer_id: 1},
		{Id: 1, Theme: "Изготовление латте", Text: "Необходима ли молочная пена при изготовлении латте", Right_answer_id: 4},
		{Id: 1, Theme: "Сиропы для латте", Text: "Какого сиропа нет в нашей сети кофеен", Right_answer_id: 6},
	}

	c.JSON(http.StatusOK, questions)
}
