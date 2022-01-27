package validation

type ErrorResponse struct {
	Error      string   `json:"error"`       // First error
	StatusCode int      `json:"status_code"` // Status Code
	Code       int      `json:"code"`        // Error code
	Errors     []string `json:"errors"`      // List of known errors
}

func NewErrorResponse() ErrorResponse {
	return ErrorResponse{}
}

func NewError(message string, statusCode int) ErrorResponse {
	// Add message to list of errors
	errors := make([]string, 0)
	errors = append(errors, message)

	// Create and return new error
	return ErrorResponse{Error: message, StatusCode: statusCode, Errors: errors}
}

func (e *ErrorResponse) AddError(message string, statusCode int, code int) *ErrorResponse {
	// Add message to list of errors
	e.Errors = append(e.Errors, message)

	// Override last error
	e.Error = message
	e.StatusCode = statusCode
	e.Code = code

	return e
}
