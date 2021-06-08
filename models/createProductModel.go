package models

type NewProduct struct {
	Name string `json:"name"` //Имя товара
	Price ProductDataInMSPrice `json:"buyPrice"` //Закупочная Цена Вложенный json
}

func NewProductModel(name string, price float64) *NewProduct {
      return &NewProduct{
		  Name: name,
		  Price: ProductDataInMSPrice{
			  Value: price,
		  },
	  }
}