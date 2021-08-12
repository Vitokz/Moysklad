package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Vitokz/Moysklad/proto"
)

type NewSupply struct { //Модель данных для создание приемки
	Name        string             `json:"name"`         //Имя приемки
	Organiztion SupplyOrganization `json:"organization"` //Ссылка на на твою организацию
	Agent       SupplyAgent        `json:"agent"`        //Контрагент=поставщик
	Store       SupplyStore        `json:"store"`        //Ссылка на твой склад
} //++

type SupplyOrganization struct { //Метаданные организации
	Meta Meta `json:"meta"`
} //++

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
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
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

func MakeNewSupply(nameSupply string, nameAgent string, id string) *NewSupply {
	agent, err := SearchAgent(nameAgent, id)

	_ = err

	//if nameAgent == "Предваритеный счет" {
	//	agent = SupplyAgent{
	//		Meta: Meta{
	//			Href:         "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/db58c2b4-c1f0-11eb-0a80-007c001a57b0",
	//			MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/metadata",
	//			Type:         "counterparty",
	//			MediaType:    "application/json",
	//			UuidHref:     "https://online.moysklad.ru/app/#company/edit?id=db58c2b4-c1f0-11eb-0a80-007c001a57b0",
	//		},
	//	}
	//}
	return &NewSupply{
		Name: nameSupply,
		Organiztion: SupplyOrganization{
			Meta: Meta{
				Href:         "https://online.moysklad.ru/api/remap/1.2/entity/organization/2ffd9caf-f840-11eb-0a80-05da0003924f",
				MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/organization/metadata",
				Type:         "organization",
				MediaType:    "application/json",
				UuidHref:     "https://online.moysklad.ru/app/#mycompany/edit?id=2ffd9caf-f840-11eb-0a80-05da0003924f",
			},
		},
		Store: SupplyStore{
			Meta: Meta{
				Href:         "https://online.moysklad.ru/api/remap/1.2/entity/store/2fff0410-f840-11eb-0a80-05da00039251",
				MetadataHref: "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
				Type:         "store",
				MediaType:    "application/json",
				UuidHref:     "https://online.moysklad.ru/app/#warehouse/edit?id=2fff0410-f840-11eb-0a80-05da00039251",
			},
		},
		Agent: agent,
	}
}
func SearchAgent(nameAgent, idUser string) (SupplyAgent, error) {
	client := &http.Client{} //Создание клиента запроса

	url := proto.COUNTERPARTY_URL
	req, err := http.NewRequest("GET", url, nil) //Создание запроса
	if err != nil {
		return SupplyAgent{}, err
	}

	q := req.URL.Query() //Добавление GET параметров
	q.Add("search", nameAgent)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+idUser)

	resp, err := client.Do(req) //Отправка запроса
	if err != nil {
		return SupplyAgent{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body) //Обработка ответа
	if err != nil {
		return SupplyAgent{}, err
	}

	rows := new(RowsAgent) //Создание модели для получения ответа от мойсклад
	json.Unmarshal([]byte(bodyText), &rows)
	if len(rows.Rows) == 0 {
		return SupplyAgent{}, fmt.Errorf("no find products")
	}
	result := rows.Rows[0]
	return result, nil //Возврат готовой модели товара из МойСклад
}

type RowsAgent struct {
	Rows []SupplyAgent `json:"rows"`
}
type SupplyResult struct {
	Id         string      `json:"id"`
	Exceptions []Exception `json:"exceptions"`
}

func TakeAliasModel(value string) Attribute {
	return Attribute{
		Meta: Meta{
			Href:      "https://online.moysklad.ru/api/remap/1.2/entity/product/metadata/attributes/ccf8afed-f851-11eb-0a80-05da0004ace2",
			Type:      "attributemetadata",
			MediaType: "application/json",
		},
		ID:    "ccf8afed-f851-11eb-0a80-05da0004ace2",
		Type:  "text",
		Name:  "Alias",
		Value: value,
	}
}

type Exception struct {
	Name   string  `json:"name"`
	Count  int     `json:"count"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func ExceptionCreate(product CsvProducts, status string) Exception {
	return Exception{
		Name:   product.Name,
		Count:  product.Count,
		Price:  product.Price,
		Status: status,
	}
}
