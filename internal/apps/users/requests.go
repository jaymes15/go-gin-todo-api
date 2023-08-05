package users

// Binding from JSON
type Register struct {
	UserName string `form:"name" json:"username" binding:"required,min=2"`
	Password string `form:"password" json:"password"  binding:"required,min=5"`
}

type Login struct {
	UserName string `form:"name" json:"username" binding:"required,min=2"`
	Password string `json:"password"  binding:"required,min=5"`
}
