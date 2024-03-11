package model

type Lesson struct {
	Id            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Passing_Score int    `json:"passing_score"`
}

type LessonResponse struct {
	Id            int    `json:"int"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Passing_Score int    `json:"passing_score"`
}
