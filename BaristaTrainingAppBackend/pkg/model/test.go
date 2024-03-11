package model

type Test struct {
	Id          int    `json:"-"`
	Theme       string `json:"theme" binding:"required"`
	Description string `json:"description" binding:"required"`
	Role_id     int    `json:"role_id" binding:"required"`
}

type TestResponse struct {
	Id          int    `json:"int"`
	Theme       string `json:"theme" binding:"required"`
	Description string `json:"description" binding:"required"`
	Role_id     int    `json:"role_id" binding:"required"`
}
