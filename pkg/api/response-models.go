package api

type ErrorModel struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
