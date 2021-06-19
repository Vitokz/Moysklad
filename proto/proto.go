package proto

const (
	NEW_FILE_XLS            = "proto/priceList/Товары.xls"
	NEW_FILE_CSV            = "proto/priceList/Товары.csv"
	RELATIONS_FILE_PATH     = "proto/relations/Relation.xlsx"
	CONFIG_REST_PATH        = "config/rest.toml"
	BASIC_AUTH_URL_MOYSKLAD = "https://online.moysklad.ru/api/remap/1.2/security/token"
	PRODUCTS_IN_MY_SKLAD    = "https://online.moysklad.ru/api/remap/1.2/entity/product"
	CREATE_SUPPLY_URL       = "https://online.moysklad.ru/api/remap/1.2/entity/supply"
	COUNTERPARTY_URL        = "https://online.moysklad.ru/api/remap/1.2/entity/counterparty"
	LOGIN                   = "admin@myskladapptest"
	PASSWORD                = "9a99f03fd9"
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
	case "Предваритеный счет":
		return CSVIDS{
			Number: 0,
			Name:   1,
			Price:  6,
			Count:  8,
			Sum:    9,
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
