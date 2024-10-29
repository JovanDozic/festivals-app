package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	// todo: will we add user_profile info here too?
}

type RegisterAttendeeRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// todo: add other user_profile info here
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
