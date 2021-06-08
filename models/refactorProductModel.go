package models

type RefactorProductsInMS struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RowsForRefactor struct { //Структура для json ответа запроса всех товаров в МС :Пока что нужно только в эндпоинте /sort
	Rows []RefactorProductsInMS `json:"rows"`
}