package response

const (
	MessageRegistration   = "registration successfull"
	MessageLogin          = "login successfull"
	MessageAuthentication = "authentication successfull"
	MessageLogout         = "logout successfull"
)

type UserResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewUserResponse(name, email, message, token string) UserResponse {
	return UserResponse{
		Name:    name,
		Email:   email,
		Message: message,
		Token:   token,
	}
}