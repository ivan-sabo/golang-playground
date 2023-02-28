package models

// Standard error reponse.
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
