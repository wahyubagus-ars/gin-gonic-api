package dto

type AddUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
}
