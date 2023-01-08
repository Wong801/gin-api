package entity

type HttpResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type ResultError struct {
	Reason string `json:"reason"`
}
