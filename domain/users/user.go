package users

type GetUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type GetUserErrorResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
