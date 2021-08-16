package ds

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"user" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
