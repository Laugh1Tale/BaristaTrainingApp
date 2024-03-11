package model

type Lection struct {
	Id    int    `json:"-"`
	Theme string `json:"theme" binding:"required"`
}

type LectionResponse struct {
	Id    int    `json:"int"`
	Theme string `json:"theme" binding:"required"`
}
