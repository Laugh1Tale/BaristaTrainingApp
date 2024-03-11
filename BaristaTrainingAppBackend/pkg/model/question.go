package model

type Question struct {
	Id              int    `json:"-"`
	Theme           string `json:"theme" binding:"required"`
	Text            string `json:"text" binding:"required"`
	Right_answer_id int    `json:"right_answer_id" binding:"required"`
}

type QuestionResponse struct {
	Id              int    `json:"int"`
	Theme           string `json:"theme" binding:"required"`
	Text            string `json:"text" binding:"required"`
	Right_answer_id int    `json:"right_answer_id" binding:"required"`
}
