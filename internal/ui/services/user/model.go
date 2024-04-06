package user_service

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Role     string `json:"role,omitempty"`
}
