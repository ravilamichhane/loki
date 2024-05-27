package common

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}
