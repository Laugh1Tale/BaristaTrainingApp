package model

type Answer struct {
	Id   int    `json:"-"`
	Text string `json:"text" binding:"required"`
}

type AnswerResponse struct {
	Id   int    `json:"int"`
	Text string `json:"text" binding:"required"`
}
