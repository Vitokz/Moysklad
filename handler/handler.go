package handler

import (
	"github.com/Vitokz/Moysklad/file"
)

type Handler struct {
	Xlsx *file.File
}

//Cоздание нового хэндлера
func NewHandler(f *file.File) *Handler {
	return &Handler{
		Xlsx: f,
	}
}
