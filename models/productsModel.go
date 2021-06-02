package models

type Product struct { // Структура для представления одного товара :Пока что нужно только в эндпоинте /sort
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type XLSXProducts struct { //Структура ассоциотивного массива из functionsFile.go
	Name string
	Keys string
}

type Rows struct { //Структура для json ответа запроса всех товаров в МС
	Rows []Product `json:"rows"`
}