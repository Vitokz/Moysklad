package file

import (
	"github.com/Vitokz/Moysklad/proto"
	"github.com/tealeg/xlsx/v3"
)

type File struct {  //Структура с открытым файлом xlsx с соотношениями 
	File *xlsx.File  //Ссылка для общения с файлом
}

//Создание структуры для работы с файлом соотношений
func NewFileOpen() *File { 
	file, err := xlsx.OpenFile(proto.RELATIONS_FILE_PATH)
	if err != nil {
		panic(err)
	}

	return &File{
		File: file,
	}
}
