package models

type NewProduct struct {
	Name       string               `json:"name"`                 //Имя товара
	Price      ProductDataInMSPrice `json:"buyPrice"`             //Закупочная Цена Вложенный json
	Attributes []Attribute          `json:"attributes,omitempty"` //Аттрибуты товара
}

func NewProductModel(name string, desc string, price float64) *NewProduct {
	return &NewProduct{
		Name: name,
		Price: ProductDataInMSPrice{
			Value: price,
		},
		Attributes: []Attribute{
			TakeAliasModel(desc),
		},
	}
}
