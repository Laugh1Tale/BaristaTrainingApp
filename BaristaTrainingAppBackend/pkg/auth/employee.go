package auth

type Employee struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
