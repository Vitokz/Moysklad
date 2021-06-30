package models

type Product struct { // Структура для представления одного товара :Пока что нужно только в эндпоинте /sort
	Meta *Meta  `json:"meta"` // Метаданные Товара
	Id   string `json:"id"`   // ID Товара (Только для чтения)
	// AccountID     string        `json:"accountId"`               // ID учетной записи (Только для чтения)
	// Owner         employee      `json:"owner"`                   // Метаданные владельца (Сотрудника)
	// Shared        bool          `json:"shared"`                  // Общий доступ
	// Group         group         `json:"group"`                   // Метаданные отдела сотрудника
	// SyncID        string        `json:"syncId,omitempty"`        // ID синхронизации
	// Updated       string        `json:"updated"`                 // Момент последнего обновления сущности (Только для чтения)
	Name        string `json:"name"`        // Наименование Товара
	Description string `json:"description"` // Описание Товара
	// Code          string        `json:"code,omitempty"`          // Код Товара
	// ExternalCode  string        `json:"externalCode"`            // Внешний код Товара
	// Archived      bool          `json:"archived"`                // Добавлен ли Товар в архив
	// PathName      string        `json:"pathName"`                // Наименование группы, в которую входит Товар (Только для чтения)
	// Vat           int           `json:"vat,omitempty"`           // НДС %
	// EffectiveVat  int           `json:"effectiveVat,omitempty"`  // Реальный НДС % (Только для чтения)
	//ProductFolder productFolder `json:"productFolder,omitempty"` // Метаданные группы Товара
	//Uom           uom           `json:"uom,omitempty"`           // Единицы измерения
	//Images        images        `json:"images,omitempty"`        // Изображения
	//MinPrice      price         `json:"minPrice,omitempty"`      // Минимальная цена
	//SalePrices    []salePrices  `json:"salePrices,omitempty"`    // Цены продажи
	BuyPrice        buyPrice           `json:"buyPrice,omitempty"` // Закупочная цена
	//Supplier      counterParty  `json:"supplier,omitempty"`      // Метаданные контрагента-поставщика
	Attributes []Attribute `json:"attributes,omitempty"` // Коллекция доп. полей
	//Country       country       `json:"country,omitempty"`       // Метаданные Страны
	//Article       string        `json:"article,omitempty"`       // Артикул
	//Weight        float64       `json:"weight,omitempty"`        // Вес
	//Volume        float64       `json:"volume,omitempty"`        // Объем
	//Packs         []struct {
	//	ID       string    `json:"id"`                 // ID упаковки товара (Только для чтения)
	//	Uom      uom       `json:"uom"`                // Метаданные единиц измерения
	//	Quantity int       `json:"quantity"`           // Количество Товаров в упаковке данного вида
	//	Barcodes []barcode `json:"barcodes,omitempty"` // Массив штрихкодов упаковок товаров
	//} `json:"packs,omitempty"` // Упаковки Товара
	//Alcoholic struct {
	// 	Excise   bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	// 	Type     int     `json:"type,omitempty"`     // Код вида продукции
	// 	Strength float64 `json:"strength,omitempty"` // Крепость
	// 	Volume   float64 `json:"volume,omitempty"`   // Объём тары
	// } `json:"alcoholic,omitempty"` // Объект, содержащий поля алкогольной продукции
	// VariantsCount      int       `json:"variantsCount"`               // Количество модификаций у данного товара (Только для чтения)
	// MinimumBalance     float64   `json:"minimumBalance"`              // Неснижаемый остаток
	// IsSerialTrackable  bool      `json:"isSerialTrackable,omitempty"` // Учет по серийным номерам. Не может быть указан вместе с alcoholic и weighed
	// Things             []string  `json:"things,omitempty"`            // Серийные номера
	// Barcodes           []barcode `json:"barcodes,omitempty"`          // Штрихкоды
	// DiscountProhibited bool      `json:"discountProhibited"`          // Признак запрета скидок
	// Tnved              string    `json:"tnved,omitempty"`             // Код ТН ВЭД
	// TrackingType       string    `json:"trackingType,omitempty"`      // Тип маркируемой продукции
	// PaymentItemType    string    `json:"paymentItemType,omitempty"`   // Признак предмета расчета
	// TaxSystem          string    `json:"taxSystem,omitempty"`         // Код системы налогообложения
	// PPEType            string    `json:"ppeType,omitempty"`           // Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13)
	// Files              struct {
	// 	Meta `json:"meta"`
	// } `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}

type XLSXProducts struct { //Структура ассоциотивного массива из functionsFile.go :Пока что нужно только в эндпоинте /sort
	Name string
	Keys string
}

type ProductRows struct { //Структура для json ответа запроса всех товаров в МС :Пока что нужно только в эндпоинте /sort
	Rows []Product `json:"rows"`
}
