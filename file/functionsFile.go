package file

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

var products = make(map[string]string, 0)  //Массив с соотношениями из 1с [имя] = [ключи]

func (f *File) SortAllNames() (map[string]string,error) {
	sh, ok := f.File.Sheet["TDSheet"] //Открываю главную страницу в файле с соотношениями
	defer sh.Close() //По окончании закрываю его
	if !ok {
		err:=fmt.Errorf("Sheet does not exist")
		return nil,err
	}
	err := sh.ForEachRow(rowVisitor)  //Прохожусь по каждой строке в таблице и передаю callback функцию
	if err != nil {
		fmt.Println("failed to sorting rows")
		return nil,err
	}
	return products,nil //Возвращаю готовый массив со всеми элементами
}

func rowVisitor(r *xlsx.Row) error {
	name := fmt.Sprintf("%v", r.GetCell(1)) //Достаю первую ячейку в строке это имя продукта
	key := fmt.Sprintf("%v;", r.GetCell(5)) //Достаю 5 ячейку в строке это ключ продукта 
	                                        //при их множественном кол-ве они соединяются в одну строку с разделением через ;
	products[name] += key
	return nil
}
