package proto

const (
	LIMIT_PRODUCTS      = 1000
	NEW_FILE_XLS        = "proto/priceList/Товары.xls"
	NEW_FILE_CSV        = "proto/priceList/Товары.csv"
	RELATIONS_FILE_PATH = "proto/relations/Relation.xlsx"
	CONFIG_REST_PATH    = "config/rest.toml"
	LOGIN               = "admin@amritasportwork"
	PASSWORD            = "2823e8e337"
)

const ( //Urls
	BASIC_AUTH_URL_MOYSKLAD = "https://online.moysklad.ru/api/remap/1.2/security/token"
	PRODUCTS_IN_MY_SKLAD    = "https://online.moysklad.ru/api/remap/1.2/entity/product"
	CREATE_SUPPLY_URL       = "https://online.moysklad.ru/api/remap/1.2/entity/supply"
	COUNTERPARTY_URL        = "https://online.moysklad.ru/api/remap/1.2/entity/counterparty"
)

const ( //convertio
	///KEY = "29bac52690fcf8e0454e5e7ca96fed55"
	KEY = "a86cdbd799a71c05f50b8a2bc556515e"
)

type CSVIDS struct {
	Number int
	Name   int
	Price  int
	Count  int
	Sum    int
}

func StatCSVFile(agent string) CSVIDS {
	switch agent {
	case "АРТ Современные научные технологии":
		return CSVIDS{
			Number: 0,
			Name:   3,
			Price:  6,
			Count:  5,
			Sum:    7,
		}
	case "БИГ Мэджик Фуд":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  11,
			Count:  8,
			Sum:    12,
		}
	case "Бомббар":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  9,
			Count:  6,
			Sum:    10,
		}
	case "Ганза Спорт":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  6,
			Count:  8,
			Sum:    10,
		}
	case "Ё-Батоны":
		return CSVIDS{
			Number: 1,
			Name:   3,
			Price:  25,
			Count:  20,
			Sum:    29,
		}
	case "Клуб Здоровья":
		return CSVIDS{
			Number: 1,
			Name:   7,
			Price:  29,
			Count:  24,
			Sum:    33,
		}
	case "Оптстронг Бурганов":
		return CSVIDS{
			Number: 0,
			Name:   3,
			Price:  39,
			Count:  36,
			Sum:    42,
		}
	case "Сойка":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  4,
			Count:  2,
			Sum:    5,
		}
	case "Спортпит-Инвест":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  6,
			Count:  8,
			Sum:    10,
		}
	case "ИП Дейкина О.Е.":
		return CSVIDS{
			Number: 1,
			Name:   5,
			Price:  60,
			Count:  50,
			Sum:    68,
		}
	case "Фитбаропт":
		return CSVIDS{
			Number: 1,
			Name:   3,
			Price:  44,
			Count:  35,
			Sum:    49,
		}
	case "Фитхаус (Чечня)":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  4,
			Count:  2,
			Sum:    5,
		}
	default:
		return CSVIDS{}
	}
}

// "meta": { //organiztion
// 	"href": "https://online.moysklad.ru/api/remap/1.2/entity/organization/db577e9f-c1f0-11eb-0a80-007c001a57ad",
// 	"metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/organization/metadata",
// 	"type": "organization",
// 	"mediaType": "application/json",
// 	"uuidHref": "https://online.moysklad.ru/app/#mycompany/edit?id=db577e9f-c1f0-11eb-0a80-007c001a57ad"
// }

// "meta": { //один из поставщмков contragent
// 	"href": "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/db58c2b4-c1f0-11eb-0a80-007c001a57b0",
// 	"metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/counterparty/metadata",
// 	"type": "counterparty",
// 	"mediaType": "application/json",
// 	"uuidHref": "https://online.moysklad.ru/app/#company/edit?id=db58c2b4-c1f0-11eb-0a80-007c001a57b0"
// }

// "meta": {  //store
// 	"href": "https://online.moysklad.ru/api/remap/1.2/entity/store/db58bba2-c1f0-11eb-0a80-007c001a57af",
// 	"metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
// 	"type": "store",
// 	"mediaType": "application/json",
// 	"uuidHref": "https://online.moysklad.ru/app/#warehouse/edit?id=db58bba2-c1f0-11eb-0a80-007c001a57af"
// }
