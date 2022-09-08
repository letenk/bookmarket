package web

type RegisterInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Address  string `json:"address" binding:"required"`
	City     string `json:"city" binding:"required"`
	Province string `json:"province" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"password"`
	Role     string `json:"role" binding:"required"`
}
