package models

// This text will appear as description of your response body.
// swagger:response ErrorResponse
type ErrorResponseWrapper struct {
	// in: body
	Body ErrorResponse
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}
