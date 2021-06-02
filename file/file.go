package file

import (
	"github.com/Vitokz/Moysklad/proto"
	"github.com/tealeg/xlsx/v3"
)

type File struct {  //Структура с открытым файлом xlsx с соотношениями 
	File *xlsx.File
}

func NewFileOpen() *File { //Создание струтуры File
	file, err := xlsx.OpenFile(proto.RELATIONS_FILE_PATH)
	if err != nil {
		panic(err)
	}

	return &File{
		File: file,
	}
}
