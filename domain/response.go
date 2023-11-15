package domain

type ErrorResponse struct {
	Error string `json:"error"`
}

type NormalResponse struct {
	Server string `json:"server"`
}
