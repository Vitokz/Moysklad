package models

type NewSupply struct { //Модель данных для создание приемки
	Name        string             `json:"name"`         //Имя приемки
	Organiztion SupplyOrganization `json:"organization"` //Ссылка на на твою организацию
	Agent       SupplyAgent        `json:"agent"`        //Контрагент=поставщик
	Store       SupplyStore        `json:"store"`        //Ссылка на твой склад
}
type SupplyOrganization struct { //Метаданные организации
	Meta Meta `json:"meta"`
}
type SupplyAgent struct { //Метаданные агента
	Meta Meta `json:"meta"`
}
type SupplyStore struct { //Метаданные склада
	Meta Meta `json:"meta"`
}
type Meta struct { //Структура метаданных
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref"`
	Type         string `json:"type"`
	MediaType    string `json:"mediaType"`
	UuidHref     string `json:"uuidHref"`
}

type Supply struct { // Структура с id свежесозданной приемки
	Id string `json:"id"`
}

type CsvProducts struct { //Структура товаров взятых из CSv файла
	Name  string //csv name
	Count int    //csv count
}

type SearchProductInMS struct { //Структура для обработки запроса поиска в rows находятся все подхдящие ответы
	Rows []ProductDataInMS `json:"rows"`
}

type ProductDataInMS struct { //Структура для получения товара из массива rows с подходящими ответами
	Count int                  `json:"count"`    //Кол-во
	Price ProductDataInMSPrice `json:"buyPrice"` //Закупочная Цена Вложенный json
	Meta  Meta                 `json:"meta"`     //Метаданные
}
type ProductDataInMSPrice struct { // Структура для взятие значени цены из закупочной цены
	Value float64 `json:"value"`
}

type SupplyPostion struct { //Структура для готового запроса на добавление позиции
	Quantity   int                     `json:"quantity"`
	Price      float64                 `json:"price"`
	Assortment SupplyPostionAssortment `json:"assortment"`
}

type SupplyPostionAssortment struct {
	Meta Meta `json:"meta"`
}

func NewSupplyPosition(m ProductDataInMS) *SupplyPostion {
	return &SupplyPostion{
		Quantity: m.Count,
		Price:    m.Price.Value,
		Assortment: SupplyPostionAssortment{
			Meta: m.Meta,
		},
	}
}

func MakeNewSupply() *NewSupply {
	return &NewSupply{
		Name: "12",
		Organiztion: SupplyOrganization{
			Meta: Meta{
				Href:         "https://online.moysklad.ru/api/remap/1.2/entity/organization/db577e9f-c1f0-11eb-0a80-007c001a57ad",
				MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/organization/metadata",
				Type:         "organization",
				MediaType:    "application/json",
				UuidHref:     "https://online.moysklad.ru/app/#mycompany/edit?id=db577e9f-c1f0-11eb-0a80-007c001a57ad",
			},
		},
		Store: SupplyStore{
			Meta: Meta{
				Href:         "https://online.moysklad.ru/api/remap/1.2/entity/store/db58bba2-c1f0-11eb-0a80-007c001a57af",
				MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
				Type:         "store",
				MediaType:    "application/json",
				UuidHref:     "https://online.moysklad.ru/app/#warehouse/edit?id=db58bba2-c1f0-11eb-0a80-007c001a57af",
			},
		},
		Agent: SupplyAgent{
			Meta: Meta{
				Href:         "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/db58c2b4-c1f0-11eb-0a80-007c001a57b0",
				MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/metadata",
				Type:         "counterparty",
				MediaType:    "application/json",
				UuidHref:     "https://online.moysklad.ru/app/#company/edit?id=db58c2b4-c1f0-11eb-0a80-007c001a57b0",
			},
		},
	}
}
