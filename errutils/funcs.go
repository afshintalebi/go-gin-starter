package errutils

func GerErrorData(status int, message string, description string) HttpErrorResponse {
	if description == "" {
		description = message
	}

	return HttpErrorResponse{
		Status:      status,
		Message:     message,
		Description: description,
	}
}
