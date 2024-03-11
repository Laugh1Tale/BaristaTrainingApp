package model

type Information struct {
	Id    int    `json:"-"`
	Theme string `json:"theme" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type InformationResponse struct {
	Id    int    `json:"int"`
	Theme string `json:"theme" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
