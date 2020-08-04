package auth

type SignUpValidator struct {
	Request SignUpRequest
	auth AuthModel `json:"-"`
}

type SignUpRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}