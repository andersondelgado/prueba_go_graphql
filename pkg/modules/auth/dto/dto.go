package dto

type InputCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InputUser struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
