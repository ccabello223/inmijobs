package dto

type ResponseDTO[T any] struct {
	Status int    `json:"status"`
	Info   string `json:"info,omitempty"`
	Data   T      `json:"data"`
}

type ErrorResponseDTO struct {
	Status  int    `json:"status"`
	Info    string `json:"info,omitempty"`
	Message string `json:"message"`
}
