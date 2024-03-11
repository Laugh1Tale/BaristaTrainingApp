package model

type Course struct {
	Id            int    `json:"-"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Passing_Score int    `json:"passing_score" binding:"required"`
}

type CourseResponse struct {
	Id            int    `json:"int"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Passing_Score int    `json:"passing_score" binding:"required"`
}
