package validacao

type ErrorResponse struct {
	Message string `json:"error"`
}

type Response struct {
	Message string `json:"message"`
}

type GenericResponse[T any] struct {
	Items []T `json:"items"`
}

type ResponseGeneric[T any] struct {
	Items T `json:"items"`
}
