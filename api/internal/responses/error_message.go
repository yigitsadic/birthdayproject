package responses

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorMessage(message string) ErrorMessage {
	return ErrorMessage{
		Message: message,
	}
}
