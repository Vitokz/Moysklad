package models

type Errors struct {
	Error []Error `json:"errors"`
}
type Error struct {
	Er           string `json:"error"`
	Parameter    string `json:"parameter"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}
