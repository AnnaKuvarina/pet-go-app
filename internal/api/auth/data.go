package auth

type UserRequestData struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}

type SignupResponse struct {
	UserID string `json:"userId"`
}

type LoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
	Username string `json:"userName"`
}