package handler

import (
	"github.com/Vitokz/Moysklad/file"
)

type Handler struct {
	Xlsx *file.File
}

func NewHandler(f *file.File) *Handler {
	return &Handler{
		Xlsx: f,
	}
}
