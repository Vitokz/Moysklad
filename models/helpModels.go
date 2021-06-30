package models

// type Meta struct {
// 	Href         string `json:"href"`                   // Ссылка на объект
// 	MetadataHref string `json:"metadataHref,omitempty"` // Ссылка на метаданные сущности (Другой вид метаданных. Присутствует не во всех сущностях)
// 	Type         string `json:"type"`                   // Тип объекта
// 	MediaType    string `json:"mediaType"`              // application/json
// 	UUUIDHref    string `json:"uuidHref,omitempty"`     // Ссылка на объект на UI. Присутствует не во всех сущностях. Может быть использована для получения uuid
// 	DownloadHref string `json:"downloadHref,omitempty"` // Ссылка на скачивание Изображения. Данный параметр указывается только в meta для Изображения у Товара или Комплекта.
// 	Size         int    `json:"size,omitempty"`         // Размер выданного списка
// 	Limit        int    `json:"limit,omitempty"`        // Максимальное количество элементов в выданном списке. Максимальное количество элементов в списке равно 1000.
// 	Offset       int    `json:"offset,omitempty"`       // Отступ в выданном списке
// 	NextHref     string `json:"nextHref,omitempty"`     // Ссылка на следующую страницу сущностей.
// 	PreviousHref string `json:"previousHref,omitempty"` // Ссылка на предыдущую страницу сущностей.
// }

// TODO: MSErrors – Структура ошибки
// Возвращаемые HTTP статусы ошибок и их значения:
// 301	Запрашиваемый ресурс находится по другому URL.
// 303	Запрашиваемый ресурс может быть найден по другому URI и должен быть найден с использоваием GET запроса
// 400	Ошибка в структуре JSON передаваемого запроса
// 401	Имя и/или пароль пользователя указаны неверно или заблокированы пользователь или аккаунт
// 403	У вас нет прав на просмотр данного объекта
// 404	Запрошенный ресурс не существует
// 405	http-метод указан неверно для запрошенного ресурса
// 409	Указанный объект используется и не может быть удален
// 410	Версия API больше не поддерживается
// 412	Не указан обязательный параметр строки запроса или поле структуры JSON
// 413	Размер запроса или количество элементов запроса превышает лимит (например, количество позиций, передаваемых в массиве positions, превышает 1000)
// 429	Превышен лимит количества запросов
// 500	При обработке запроса возникла непредвиденная ошибка
// 502	Сервис временно недоступен
// 503	Сервис временно отключен
// 504	Превышен таймаут обращения к сервису, повторите попытку позднее
// type MSErrors struct {
// 	Errors []struct {
// 		Error     string `json:"error,omitempty"`         // Заголовок ошибки
// 		Parameter string `json:"parameter,omitempty"`     // Параметр, на котором произошла ошибка
// 		Code      int    `json:"code,omitempty"`          // Код ошибки (Если поле ничего не содержит, смотрите HTTP status code)
// 		Message   string `json:"error_message,omitempty"` // Сообщение, прилагаемое к ошибке.
// 	} `json:"errors"`
// }

// Group Отдел
type group struct {
	Meta *Meta  `json:"meta,omitempty"` // Метаданные Отдела
	Name string `json:"name,omitempty"` // Наименование Отдела
}

// Image структура изображения (image)
type image struct {
	Meta      *Meta  `json:"meta"`                // Метаданные объекта
	Title     string `json:"title,omitempty"`     // Название Изображения
	Filename  string `json:"filename,omitempty"`  // Имя файла
	Size      int    `json:"size,omitempty"`      // Размер файла в байтах
	Updated   string `json:"updated,omitempty"`   // Время последнего изменения
	Miniature Meta   `json:"miniature,omitempty"` // Метаданные миниатюры изображения
	Tiny      Meta   `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения
}

// Images ...
type images struct {
	Meta *Meta   `json:"meta"`
	Rows []image `json:"rows,omitempty"`
}

// SalePrices Цена продажи
type salePrices struct {
	Value     float64   `json:"value,omitempty"`     // Значение цены
	Currency  *currency `json:"currency,omitempty"`  // Ссылка на валюту в формате
	PriceType priceType `json:"priceType,omitempty"` // Тип цены
}

// ProductFolder Группа Товаров
type productFolder struct {
	Meta          *Meta          `json:"meta,omitempty"`          // Метаданные Группы Товара (Только для чтения)
	ID            string         `json:"id,omitempty"`            // ID Группы товаров (Только для чтения)
	AccountID     string         `json:"accountId,omitempty"`     // ID учетной записи (Только для чтения)
	Owner         *employee      `json:"owner,omitempty"`         // Метаданные владельца (Сотрудника)
	Shared        bool           `json:"shared,omitempty"`        // Общий доступ
	Group         *group         `json:"group,omitempty"`         // Метаданные отдела сотрудника
	Updated       string         `json:"updated,omitempty"`       // Момент последнего обновления сущности (Только для чтения)
	Name          string         `json:"name,omitempty"`          // Наименование Группы товаров
	Description   string         `json:"description,omitempty"`   // Описание Группы товаров
	Code          string         `json:"Description,omitempty"`   // Код Группы товаров
	ExternalCode  string         `json:"externalCode,omitempty"`  // Внешний код Группы товаров
	Archived      bool           `json:"archived,omitempty"`      // Добавлена ли Группа товаров в архив (Только для чтения)
	PathName      string         `json:"pathName,omitempty"`      // Наименование Группы товаров, в которую входит данная Группа товаров (Только для чтения)
	Vat           int            `json:"vat,omitempty"`           // НДС %
	EffectiveVat  int            `json:"effectiveVat,omitempty"`  // Реальный НДС % (Только для чтения)
	ProductFolder *productFolder `json:"productFolder,omitempty"` // Ссылка на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных
	TaxSystem     string         `json:"taxSystem,omitempty"`     // Код системы налогообложения
	// --
	// taxSystem
	// TAX_SYSTEM_SAME_AS_GROUP	Совпадает с группой
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --
}

// Currency Валюта
type currency struct {
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные Валюты
	ID             string        `json:"id,omitempty"`             // ID Валюты (Только для чтения)
	Name           string        `json:"name,omitempty"`           // Краткое аименование Валюты
	FullName       string        `json:"fullName,omitempty"`       // Полное наименование Валюты
	Code           string        `json:"code,omitempty"`           // Цифровой код Валюты
	ISOCode        string        `json:"isoCode,omitempty"`        // Буквенный код Валюты
	Rate           float64       `json:"rate,omitempty"`           // Курс Валюты
	Multiplicity   int           `json:"multiplicity,omitempty"`   // Кратность курса Валюты
	Indirect       bool          `json:"indirect,omitempty"`       // Признак обратного курса Валюты
	RateUpdateType bool          `json:"rateUpdateType,omitempty"` // Способ обновления курса Валюты (Только для чтения)
	MajorUnit      *currencyUnit `json:"majorUnit,omitempty"`      // Формы единиц целой части Валюты
	MinorUnit      *currencyUnit `json:"minorUnit,omitempty"`      // Формы единиц дробной части Валюты
	Archived       bool          `json:"archived,omitempty"`       // Добавлена ли Валюта в архив
	System         bool          `json:"system,omitempty"`         // Основана ли валюта на валюте из системного справочника (Только для чтения)
	Default        bool          `json:"default,omitempty"`        // Является ли валюта валютой учета (Только для чтения)
}

// CurrencyUnit Форма единиц
type currencyUnit struct {
	Meta   *Meta  `json:"meta,omitempty"`   // Метаданные Формы единиц
	Gender string `json:"gender,omitempty"` // Грамматический род единицы валюты (допустимые значения masculine - мужской, feminine - женский)
	S1     string `json:"s1,omitempty"`     // Форма единицы, используемая при числительном 1
	S2     string `json:"s2,omitempty"`     // Форма единицы, используемая при числительном 2
	S5     string `json:"s5,omitempty"`     // Форма единицы, используемая при числительном 5
}

// Rate ...
type rate struct {
	Currency *currency `json:"currency,omitempty"`
}

// BuyPrice Закупочная цена
type buyPrice struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *currency `json:"currency,omitempty"` // Ссылка на валюту в формате
}

// Attribute Дополнительное поле
type Attribute struct {
	Meta  Meta   `json:"meta"`  // Ссылка на метаданные доп. поля
	ID    string `json:"id"`    // Id соответствующего доп. поля
	Type  string `json:"type"`  // Тип доп. поля
	Name  string `json:"name"`  // Наименование доп. поля
	Value string `json:"value"` // Значение, указанное в доп. поле.
}

// Barcode Штрихкоды
type barcode struct {
	EAN13   string `json:"ean13,omitempty"`   // штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
	EAN8    string `json:"ean8,omitempty"`    // штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
	Code128 string `json:"code128,omitempty"` // штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
	GTIN    string `json:"gtin,omitempty"`    // штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
	UPC     string `json:"upc,omitempty"`     // штрихкод в формате UPC, если требуется создать штрихкод в формате UPC.
}

// File Файл
type file struct {
	Meta      Meta     `json:"meta"`                // Метаданные объекта
	Title     string   `json:"title,omitempty"`     // Название Файла
	Filename  string   `json:"filename,omitempty"`  // Имя Файла
	Size      int      `json:"size,omitempty"`      // Размер Файла в байтах
	Created   string   `json:"created,omitempty"`   // Время загрузки Файла на сервер
	CreatedBy employee `json:"createdBy,omitempty"` // Метаданные сотрудника, загрузившего Файл
	Miniature Meta     `json:"miniature,omitempty"` // Метаданные миниатюры изображения (поле передается только для Файлов изображений)
	Tiny      Meta     `json:"tiny,omitempty"`      //	Метаданные уменьшенного изображения (поле передается только для Файлов изображений)
}

// Characteristics ...
type characteristics []struct {
	Meta  *Meta  `json:"meta,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Uom Единица измерения
type uom struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные Единицы измерения
}

// Stock ...
type stock struct {
	Stock map[string]string `json:"stock,omitempty"`
}

// Country Метаданные страны
type country struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные страны
}

// Region Метаданные региона
type region struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные региона
}

// AddressFull Адрес
type addressFull struct {
	PostalCode string   `json:"postalCode,omitempty"` // Почтовый индекс
	Country    *country `json:"country,omitempty"`    // Метаданные страны
	Region     *region  `json:"region,omitempty"`     // Метаданные региона
	City       string   `json:"city,omitempty"`       // Город
	Street     string   `json:"street,omitempty"`     // Улица
	House      string   `json:"house,omitempty"`      // Дом
	Apartment  string   `json:"apartment,omitempty"`  // Квартира
	AddInfo    string   `json:"addInfo,omitempty"`    // Другое
	Comment    string   `json:"comment,omitempty"`    // Комментарий
}

// LastOperation Последние операции
type lastOperation struct {
	Entity string `json:"entity,omitempty"` // Ключевое слово, обозначающее тип последней операции (Только для чтения)
	Name   string `json:"name,omitempty"`   // Наименование (номер) последней операции (Только для чтения)
}

// Status Статус документа
type status struct {
	Meta      *Meta  `json:"meta,omitempty"`      // Метаданные Статуса
	ID        string `json:"id,omitempty"`        // ID Статуса (Только для чтения)
	AccountID string `json:"accountId,omitempty"` // ID учетной записи (Только для чтения)
	Name      string `json:"name,omitempty"`      // Наименование Статуса
	Color     string `json:"color,omitempty"`     // Цвет Статуса
	StateType string `json:"stateType,omitempty"` // Тип Статуса
	// --
	// Regular	Обычный (значение по умолчанию)
	// Successful	Финальный положительный
	// Unsuccessful	Финальный отрицательный
	// --
	EntityType string `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (ключевое слово в рамках JSON API)
}

// PrintTemplate Шаблон печатной формы
type printTemplate struct {
	Meta    *Meta  `json:"meta,omitempty"`    // 	Метаданные шаблона
	ID      string `json:"id,omitempty"`      // ID шаблона
	Name    string `json:"name,omitempty"`    // Наименование шаблона
	Type    string `json:"type,omitempty"`    // Тип шаблона (entity - документ)
	Content string `json:"content,omitempty"` //	Ссылка на скачивание
}
type employee struct {
	Meta         *Meta     `json:"meta"`                   // Метаданные Сотрудника
	ID           string    `json:"id,omitempty"`           // ID Сотрудника (Только для чтения)
	AccountID    string    `json:"accountId,omitempty"`    // ID учетной записи (Только для чтения)
	Owner        *employee `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       bool      `json:"shared,omitempty"`       // Общий доступ
	Group        *group    `json:"group,omitempty"`        // Отдел сотрудника
	Updated      string    `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника (Только для чтения)
	Name         string    `json:"name,omitempty"`         // Наименование Сотрудника (Только для чтения)
	Description  string    `json:"description,omitempty"`  // Комментарий к Сотруднику
	ExternalCode string    `json:"externalCode,omitempty"` // Внешний код Сотрудника (Только для чтения)
	Archived     bool      `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
	Created      string    `json:"created,omitempty"`      // Момент создания Сотрудника (Только для чтения)
	UID          string    `json:"uid,omitempty"`          // Логин Сотрудника (Только для чтения)
	Email        string    `json:"email,omitempty"`        // Электронная почта сотрудника
	Phone        string    `json:"phone,omitempty"`        // Телефон сотрудника
	FirstName    string    `json:"firstName,omitempty"`    // Имя
	MiddleName   string    `json:"middleName,omitempty"`   // Отчество
	LastName     string    `json:"lastName,omitempty"`     // Фамилия
	FullName     string    `json:"fullName,omitempty"`     // Имя Отчество Фамилия (Только для чтения)
	ShortFio     string    `json:"shortFio,omitempty"`     // Краткое ФИО (Только для чтения)
	//Cashiers     []Cashier   `json:"cashiers,omitempty"`     // Массив кассиров (Только для чтения)
	Attributes []Attribute `json:"attributes,omitempty"` // Дополнительные поля Сотрудника // TODO
	Image      *image      `json:"image,omitempty"`      // Фотография сотрудника
	INN        string      `json:"inn,omitempty"`        // ИНН сотрудника (в формате ИНН физического лица)
	Position   string      `json:"position,omitempty"`   // Должность сотрудника
}
type priceType struct {
	Meta         *Meta  `json:"meta,omitempty"`         // Метаданные Типа цены (Только для чтения)
	ID           string `json:"id,omitempty"`           // ID типа цены (Только для чтения)
	Name         string `json:"name,omitempty"`         // Наименование Типа цены
	ExternalCode string `json:"externalCode,omitempty"` // Внешний код Типа цены
}
type price struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}
type counterParty struct {
	Meta               *Meta        `json:"meta"`                         // Метаданные Контрагента
	ID                 string       `json:"id,omitempty"`                 // ID Контрагента Только для чтения
	AccountID          string       `json:"accountId,omitempty"`          // ID учетной записи Только для чтения
	Owner              *employee    `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Shared             bool         `json:"shared,omitempty"`             // Общий доступ
	Group              group        `json:"group,omitempty"`              // Отдел сотрудника
	SyncID             string       `json:"syncId,omitempty"`             // ID синхронизации После заполнения недоступен для изменения
	Updated            string       `json:"updated,omitempty"`            // Момент последнего обновления Контрагента Только для чтения
	Name               string       `json:"name,omitempty"`               // Наименование Контрагента Необходимое при создании
	Description        string       `json:"description,omitempty"`        // Комментарий к Контрагенту
	Code               string       `json:"code,omitempty"`               // Код Контрагента
	ExternalCode       string       `json:"externalCode,omitempty"`       // Внешний код Контрагента Только для чтения
	Archived           bool         `json:"archived,omitempty"`           // Добавлен ли Контрагент в архив
	Created            string       `json:"created,omitempty"`            // Момент создания
	Email              string       `json:"email,omitempty"`              // Адрес электронной почты
	Phone              string       `json:"phone,omitempty"`              // Номер городского телефона
	FAX                string       `json:"fax,omitempty"`                // Номер факса
	ActualAddress      string       `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	ActualAddressFull  addressFull  `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Accounts           accounts     `json:"accounts,omitempty"`           // Массив счетов Контрагентов
	CompanyType        string       `json:"companyType,omitempty"`        // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	DiscountCardNumber string       `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	State              status       `json:"state,omitempty"`              // Метаданные Статуса Контрагента
	SalesAmount        float64      `json:"salesAmount,omitempty"`        // Сумма продаж Только для чтения
	BonusProgram       bonusProgram `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	BonusPoints        int          `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе Только для чтения
	Files              struct {
		Meta `json:"meta"`
	} `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}
type account struct {
	Meta                 *Meta         `json:"meta,omitempty"`
	ID                   string        `json:"id,omitempty"`                   // ID Счета (Только для чтения)
	AccountID            string        `json:"accountId,omitempty"`            // ID учетной записи (Только для чтения)
	Updated              string        `json:"updated,omitempty"`              // Момент последнего обновления юрлица (Только для чтения)
	IsDefault            bool          `json:"isDefault,omitempty"`            // Является ли счет основным счетом юрлица
	AccountNumber        string        `json:"accountNumber,omitempty"`        // Номер счета	Необходимое при создании
	BankName             string        `json:"bankName,omitempty"`             // Наименование банка
	BankLocation         string        `json:"bankLocation,omitempty"`         // Адрес банка
	CorrespondentAccount string        `json:"correspondentAccount,omitempty"` // Корр счет
	BIC                  string        `json:"bic,omitempty"`                  // БИК
	Agent                *organization `json:"agent,omitempty"`                // Метаданные юрлица
}

// Accounts ...
type accounts struct {
	Meta *Meta     `json:"meta,omitempty"`
	Rows []account `json:"rows,omitempty"`
}
type organization struct {
	Meta              *Meta        `json:"meta,omitempty"`              // Метаданные Юрлица
	ID                string       `json:"id,omitempty"`                // ID Юрлица (Только для чтения)
	AccountID         string       `json:"accountId,omitempty"`         // ID учетной записи (Только для чтения)
	Owner             *employee    `json:"owner,omitempty"`             // Владелец (Сотрудник)
	Shared            bool         `json:"shared,omitempty"`            // Общий доступ
	Group             *group       `json:"group,omitempty"`             // Отдел сотрудника
	SyncID            string       `json:"syncId,omitempty"`            // ID синхронизации
	Updated           string       `json:"updated,omitempty"`           // Момент последнего обновления Юрлица (Только для чтения)
	Name              string       `json:"name,omitempty"`              // Наименование Юрлица
	Description       string       `json:"description,omitempty"`       // Комментарий к Юрлицу
	Code              string       `json:"code,omitempty"`              // Код Юрлица
	ExternalCode      string       `json:"externalCode,omitempty"`      // Внешний код Юрлица (Только для чтения)
	Achived           bool         `json:"archived,omitempty"`          // Добавлено ли Юрлицо в архив
	Created           string       `json:"created,omitempty"`           // Дата создания
	ActualAddress     string       `json:"actualAddress,omitempty"`     // Фактический адрес Юрлица
	ActualAddressFull *addressFull `json:"actualAddressFull,omitempty"` // Фактический адрес Юрлица с детализацией по отдельным полям
	CompanyType       string       `json:"companyType,omitempty"`       // Тип Юрлица. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	// --
	// companyType
	// legal	Юридическое лицо
	// entrepreneur	Индивидуальный предприниматель
	// individual	Физическое лицо
	// --
	TrackingContractNumber string        `json:"trackingContractNumber,omitempty"` // Номер договора с ЦРПТ
	TrackingContractDate   string        `json:"trackingContractDate,omitempty"`   // Дата договора с ЦРПТ
	BonusProgram           *bonusProgram `json:"bonusProgram,omitempty"`           // Метаданные активной бонусной программы
	BonusPoints            int           `json:"bonusPoints,omitempty"`            // Бонусные баллы по активной бонусной программе
	LegalTitle             string        `json:"legalTitle,omitempty"`             // Полное наименование. Игнорируется, если передано одно из значений для ФИО. Формируется автоматически на основе получаемых ФИО Юрлица
	LegalLastName          string        `json:"legalLastName,omitempty"`          // Фамилия для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalFirstName         string        `json:"legalFirstName,omitempty"`         // Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalMiddleName        string        `json:"legalMiddleName,omitempty"`        // Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalAddress           string        `json:"legalAddress,omitempty"`           // Юридический адреса Юрлица
	LegalAddressFull       *addressFull  `json:"legalAddressFull,omitempty"`       // Юридический адрес Юрлица с детализацией по отдельным полям
	INN                    string        `json:"inn,omitempty"`                    // ИНН
	KPP                    string        `json:"kpp,omitempty"`                    // КПП
	OGRN                   string        `json:"ogrn,omitempty"`                   // ОГРН
	OGRNIP                 string        `json:"ogrnip,omitempty"`                 // ОГРНИП
	OKPO                   string        `json:"okpo,omitempty"`                   // ОКПО
	// certificateNumber	Meta	// Номер свидетельства
	CertificateDate string      `json:"certificateDate,omitempty"` // Дата свидетельства
	Email           string      `json:"email,omitempty"`           // Адрес электронной почты
	Phone           string      `json:"phone,omitempty"`           // Номер городского телефона
	Fax             string      `json:"fax,omitempty"`             // Номер факса
	Accounts        []account   `json:"accounts,omitempty"`        // Метаданные счетов юрлица
	Attributes      []Attribute `json:"attributes,omitempty"`      // Массив метаданных дополнительных полей юрлица
	IsEgaisEnable   bool        `json:"isEgaisEnable,omitempty"`   // Включен ли ЕГАИС для данного юрлица
	FSRARID         string      `json:"fsrarId,omitempty"`         // Идентификатор в ФСРАР
	PayerVat        bool        `json:"payerVat,omitempty"`        // Является ли данное юрлицо плательщиком НДС
	UTMURL          string      `json:"utmUrl,omitempty"`          // IP-адрес УТМ
	Director        string      `json:"director,omitempty"`        // Руководитель
	ChiefAccountant string      `json:"chiefAccountant,omitempty"` // Главный бухгалтер
}
type bonusProgram struct {
	Meta                    *Meta    `json:"meta"`                              // Метаданные Бонусной программы
	ID                      string   `json:"id"`                                // ID Бонусной программы (Только для чтения)
	AccountID               string   `json:"accountId"`                         // ID учетной записи (Только для чтения)
	Name                    string   `json:"name,omitempty"`                    // Наименование Бонусной программы
	Active                  bool     `json:"active"`                            // Индикатор, является ли бонусная программа активной на данный момент
	AllProducts             bool     `json:"allProducts"`                       // Индикатор, действует ли бонусная программа на все товары (всегда true)
	AllAgents               bool     `json:"allAgents"`                         // Индикатор, действует ли скидка на всех контрагентов
	AgentTags               []string `json:"agentTags"`                         // Тэги контрагентов, к которым применяется бонусная программа. В случае пустого значения контрагентов в результате выводится пустой массив
	EarnRateRoublesToPoint  int      `json:"earnRateRoublesToPoint,omitempty"`  // Курс начисления
	SpendRatePointsToRouble int      `json:"spendRatePointsToRouble,omitempty"` // Курс списания
	MaxPaidRatePercents     int      `json:"maxPaidRatePercents,omitempty"`     // Максимальный процент оплаты баллами
}
