package response

const (
	LevelBackend  = "backend"
	LevelDatabase = "database"
)

const (
	MessageBadRequest       = "bad request"
	MessageUnauthorized     = "unauthorized"
	MessageMethodNotAllowed = "method not allowed"
)

type ErrorResponse struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func NewErrorResponse(level, message string) ErrorResponse {
	return ErrorResponse{
		Level:   level,
		Message: message,
	}
}