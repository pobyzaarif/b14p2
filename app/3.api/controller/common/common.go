package common

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorMessage(msg string) ErrorResponse {
	return ErrorResponse{
		Message: msg,
	}
}
