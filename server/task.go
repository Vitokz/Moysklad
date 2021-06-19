package server

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Vitokz/Moysklad/models"
	"github.com/Vitokz/Moysklad/proto"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

//Эндпоинт /
func (r *Rest) GetTask(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Test endpoint",
	}).Println()
	r.Logger.WithFields(logrus.Fields{
		"event": c.FormValue("name"),
	}).Println()
	return c.JSON(http.StatusOK, r.Token)
}

//Эндпоинт /auth
func (r *Rest) Auth(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Authorization",
	}).Println()

	client := &http.Client{} //Создание клиента запроса

	req, err := http.NewRequest("POST", proto.BASIC_AUTH_URL_MOYSKLAD, nil) // Создание самого запроса
	if err != nil {
		r.Logger.WithError(err).Error()
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(proto.LOGIN, proto.PASSWORD)) //Header авторизации базовым способом
	resp, err := client.Do(req)                                                      //ЗАпуск запроса
	if err != nil {
		r.Logger.Errorf("Failed processing add login request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	bodyText, err := ioutil.ReadAll(resp.Body) //Ответ приходит в виде []byte по-этому его приходится форматировать
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = checkError(bodyText)
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	token := models.Token{}

	err = json.Unmarshal([]byte(bodyText), &token) //Перезаписываю []byte в JSON
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	r.Logger.WithFields(logrus.Fields{
		"statusAuth": "OK",
	}).Println()
	r.Token = &token
	return c.JSON(http.StatusOK, r.Token)
}

func basicAuth(username, password string) string { // Функция относящаяся к auth
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func checkError(text []byte) error {
	erro := models.Errors{}
	err := json.Unmarshal([]byte(text), &erro)
	_ = err
	if count := len(erro.Error); count != 0 {
		return fmt.Errorf("Failed : %s\n", erro.Error[0].Er)
	}
	return nil
}

//Эндпоинт /Sort
func (r *Rest) AddDescription(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start refactore description",
	}).Println()
	products, err := r.Handler.Xlsx.SortAllNames() //Достаю из своего файла с соотношениями структура имен и ключей
	if err != nil {
		r.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, "not ok")
	}
	productsXlsx := reformatXlsxProductsInStruct(products) //Рефактор из map в Struct
	_ = productsXlsx
	productsMySklad, err := r.takeProducts() //Достаю все товары из Мс
	if err != nil {
		r.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, "not ok")
	}
	err = r.checkingProducts(productsXlsx, productsMySklad) //Начинаю проверку каждого элемента на наличие в МС
	if err != nil {
		r.Logger.Errorf("Failed processing add to description: %s\n", err)
		return c.JSON(http.StatusInternalServerError, "not ok")
	}
	return c.JSON(http.StatusOK, "All requests are passed")
}

func (r *Rest) takeProducts() (*models.Rows, error) {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start to taking products MS",
		"token": r.Token.Access_token,
	})
	client := &http.Client{}

	req, err := http.NewRequest("GET", proto.PRODUCTS_IN_MY_SKLAD, nil)
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+r.Token.Access_token)
	resp, err := client.Do(req)
	if err != nil {
		r.Logger.Errorf("Failed processing take products request: %s\n", err)
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Logger.WithError(err).Error()
		return nil, err
	}
	err = checkError(bodyText)
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return nil, err
	}
	result := new(models.Rows)
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	json.Unmarshal([]byte(bodyText), &result)
	return result, nil
}

func reformatXlsxProductsInStruct(arr map[string]string) *[]models.XLSXProducts { // Функция относящаяся к sort
	result := make([]models.XLSXProducts, 0)

	for i, v := range arr {
		str := models.XLSXProducts{
			Name: i,
			Keys: strings.TrimSuffix(v, ";"),
		}
		result = append(result, str)
	}
	return &result
}

func (r *Rest) checkingProducts(XL *[]models.XLSXProducts, MS *models.Rows) error { // Функция относящаяся к sort
	r.Logger.WithFields(logrus.Fields{
		"event": "start to find identical in products",
	})
	for _, Xv := range *XL {
		for _, Mv := range MS.Rows {
			if Xv.Name == Mv.Name {
				Mv.Description = Xv.Keys
				err := r.requestToPut(Mv)
				if err != nil {
					return err
				}
			}
		}
	}
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return nil
}

func (r *Rest) requestToPut(m models.Product) error { // Функция относящаяся к sort
	client := &http.Client{} //Создание клиента запроса

	mes, err := json.Marshal(m) //Первод в []byte модели товара
	if err != nil {
		r.Logger.WithError(err).Error("Failed to encoding to json")
		return err
	}

	url := proto.PRODUCTS_IN_MY_SKLAD + "/" + m.Id                //Создание урла запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(mes)) //Сам запрос
	if err != nil {
		r.Logger.Errorf("Request fail: %s\n", err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+r.Token.Access_token)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		r.Logger.Errorf("Failed processing take products request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Logger.WithError(err).Error()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = checkError(bodyText)
	if err != nil {
		r.Logger.Errorf("Failed : %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}

//Эндпоинт /makeSupply
func (r *Rest) MakeSupply(c echo.Context) error {

	r.Logger.WithFields(logrus.Fields{
		"event": "Start creating Supply", // Добавить проверку на входные данные
	}).Print()

	nameSupply := c.FormValue("name")
	if nameSupply == "" {
		r.Logger.Error(fmt.Errorf("поле nameSupply пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	nameAgent := c.FormValue("agent")
	if nameAgent == "" {
		r.Logger.Error(fmt.Errorf("поле nameAgent пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	newFile, err := c.FormFile("file")
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	fileName, err := createFile(*newFile)
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	supply, err := makeNewSupplyInMS(r.Token.Access_token, nameSupply, nameAgent) //Создание новой приемки в мойсклад
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	file, err := openCSV(fileName) //Открытие csv файла поставщика
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	products, err := CSVRows(file, nameAgent) //Парсинг файла(Достает все строки с товарами и берет нужную информацию)
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	err, exceptions := findProductsInMs(*products, r.Token.Access_token, supply.Id) //Поиск продукта в мойсклад и в случае успеха добавление в приемку
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}

	_ = exceptions
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	})
	result := models.SupplyResult{
		Id:         supply.Id,
		Exceptions: exceptions,
	}
	os.Remove(proto.NEW_FILE_XLS)
	os.Remove(proto.NEW_FILE_CSV)
	return c.JSON(http.StatusOK, result)
}

func makeNewSupplyInMS(id string, nameSupply string, nameAgent string) (models.Supply, error) {
	supplyData := models.MakeNewSupply(nameSupply, nameAgent, id) //Создание модели с данными приемки (на данный момент все заполнено статически)
	client := &http.Client{}                                      //Создание клиента запроса
	mes, err := json.Marshal(supplyData)                          //Первод в []byte модели приемки
	if err != nil {
		return models.Supply{}, err
	}

	req, err := http.NewRequest("POST", proto.CREATE_SUPPLY_URL, bytes.NewBuffer(mes)) //Запрос на создание приемки
	if err != nil {
		return models.Supply{}, err
	}
	req.Header.Add("Authorization", "Bearer "+id)                     //Header авторизации
	req.Header.Set("Content-Type", "application/json; charset=UTF-8") // Header типа данных

	resp, err := client.Do(req)
	if err != nil {
		return models.Supply{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Supply{}, err
	}
	err = checkError(bodyText)
	if err != nil {
		return models.Supply{}, err
	}
	result := new(models.Supply) // Модель Supply Содержит в себе только id приемки
	json.Unmarshal([]byte(bodyText), result)
	return *result, nil
} //+

func CSVRows(file *os.File, agent string) (*[]models.CsvProducts, error) {
	defer file.Close() //Закрытие файла по окончании ф-ции
	//Для предварительного счета поля: 0-номер,1-имя,6-цена,8-кол-во,9-вид ед измерения,10-сумма
	result := make([]models.CsvProducts, 0) //Создание массива моделей товаров поставщика csv

	numbers := proto.StatCSVFile(agent)
	reader := csv.NewReader(file) //Чтение файла
	for {
		record, e := reader.Read() //Беру одну строку
		if e != nil {

			break
		}
		tp, _ := strconv.Atoi(record[numbers.Number]) //Проверяю первый символ строки на наличие числа ,так как это отлич знак товара
		if tp == 0 {
			continue
		} else {
			count, err := strconv.ParseFloat(record[numbers.Count], 64)
			if err != nil {
				return nil, err
			}
			price, err := strconv.ParseFloat(record[numbers.Price], 64)
			if err != nil {
				return nil, err
			}
			result = append(result, models.CsvProducts{
				Count: int(count),
				Name:  record[numbers.Name],
				Price: price,
			})
		}
	}
	return &result, nil
} //+
func findProductsInMs(products []models.CsvProducts, idUser string, idSupply string) (error, []models.CsvProducts) {

	exception := make([]models.CsvProducts, 0)

	for _, v := range products {
		check, err, ok := searchProduct(v, idUser) //Поиск продукта в списке товаров в Мойсклад
		if err != nil {
			return err, exception
		}
		if ok {
			err := addPositionInSupply(check, idSupply, idUser) //Добавление продукта в приемку
			if err != nil {
				return err, exception
			}
		} else {
			exception = append(exception, v)
		}
	}
	return nil, exception
}

func searchProduct(product models.CsvProducts, idUser string) (models.ProductDataInMS, error, bool) {
	client := &http.Client{} //Создание клиента запроса

	url := proto.PRODUCTS_IN_MY_SKLAD
	req, err := http.NewRequest("GET", url, nil) //Создание запроса
	if err != nil {
		return models.ProductDataInMS{}, err, false
	}

	q := req.URL.Query() //Добавление GET параметров
	q.Add("filter", "description~"+product.Name)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+idUser)

	resp, err := client.Do(req) //Отправка запроса
	if err != nil {
		return models.ProductDataInMS{}, err, false
	}

	bodyText, err := ioutil.ReadAll(resp.Body) //Обработка ответа
	if err != nil {
		return models.ProductDataInMS{}, err, false
	}

	err = checkError(bodyText)
	if err != nil {
		return models.ProductDataInMS{}, err, false
	}

	rows := new(models.SearchProductInMS) //Создание модели для получения ответа от мойсклад
	json.Unmarshal([]byte(bodyText), &rows)
	if len(rows.Rows) == 0 {
		return models.ProductDataInMS{}, err, false
	}
	result := rows.Rows[0]
	result.Count = product.Count
	return result, nil, true //Возврат готовой модели товара из МойСклад
} //+

func addPositionInSupply(product models.ProductDataInMS, idSupply string, idUser string) error {
	position := models.NewSupplyPosition(product) //Создание модели готовой для создания новой позиции в Мойсклад

	mes, err := json.Marshal(position) //Кодирование
	if err != nil {
		return err
	}
	url := proto.CREATE_SUPPLY_URL + "/" + idSupply + "/positions"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mes))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+idUser)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = checkError(bodyText)
	if err != nil {
		return err
	}
	return nil
}

func createFile(file multipart.FileHeader) (string, error) {
	err := createFileInput(file)
	if err != nil {
		return "", err
	}

	name, err := createFileCsv()
	if err != nil {
		return "", err
	}
	return name, nil
}

func openCSV(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Эндпоинт /makeProduct
func (r *Rest) MakeProduct(name string, price float64, desc string, c echo.Context) (models.ProductDataInMS, error) {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start creating Product",
	}).Print()

	product := models.NewProductModel(name, desc, price)

	new, err := addProductInMS(product, r.Token.Access_token)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return new, nil
}

func addProductInMS(product *models.NewProduct, id string) (models.ProductDataInMS, error) {
	new := models.ProductDataInMS{}

	mes, err := json.Marshal(product) //Кодирование
	if err != nil {
		return models.ProductDataInMS{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", proto.PRODUCTS_IN_MY_SKLAD, bytes.NewBuffer(mes))
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	req.Header.Add("Authorization", "Bearer "+id)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	err = checkError(bodyText)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	err = json.Unmarshal([]byte(bodyText), &new)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	return new, nil
}

// Эндпоинт /refactorProduct
func (r *Rest) RefactorProduct(nomecl, name string, c echo.Context) (models.ProductDataInMS, error) {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start refactoring Product",
	}).Print()

	product, err := searchProductForRefactor(nomecl, r.Token.Access_token)
	if err != nil {
		return models.ProductDataInMS{}, err
	}

	product.Description += ";" + name

	refactor, err := refactorProduct(product, r.Token.Access_token)
	if err != nil {
		return models.ProductDataInMS{}, err
	}

	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return refactor, nil
}

func searchProductForRefactor(productName string, idUser string) (models.RefactorProductsInMS, error) {
	client := &http.Client{} //Создание клиента запроса

	url := proto.PRODUCTS_IN_MY_SKLAD
	req, err := http.NewRequest("GET", url, nil) //Создание запроса
	if err != nil {
		return models.RefactorProductsInMS{}, err
	}

	q := req.URL.Query() //Добавление GET параметров
	q.Add("search", productName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+idUser)

	resp, err := client.Do(req) //Отправка запроса
	if err != nil {
		return models.RefactorProductsInMS{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body) //Обработка ответа
	if err != nil {
		return models.RefactorProductsInMS{}, err
	}
	err = checkError(bodyText)
	if err != nil {
		return models.RefactorProductsInMS{}, err
	}

	rows := new(models.RowsForRefactor) //Создание модели для получения ответа от мойсклад
	json.Unmarshal([]byte(bodyText), &rows)
	if len(rows.Rows) == 0 {
		return models.RefactorProductsInMS{}, fmt.Errorf("no find products")
	}
	result := rows.Rows[0]
	return result, nil //Возврат готовой модели товара из МойСклад
}

func refactorProduct(product models.RefactorProductsInMS, id string) (models.ProductDataInMS, error) {
	refactor := models.ProductDataInMS{}
	client := &http.Client{} //Создание клиента запроса

	mes, err := json.Marshal(product) //Первод в []byte модели товара
	if err != nil {
		return models.ProductDataInMS{}, err
	}

	url := proto.PRODUCTS_IN_MY_SKLAD + "/" + product.Id          //Создание урла запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(mes)) //Сам запрос
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	req.Header.Add("Authorization", "Bearer "+id)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	err = checkError(bodyText)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	err = json.Unmarshal([]byte(bodyText), &refactor)
	if err != nil {
		return models.ProductDataInMS{}, err
	}
	return refactor, nil
}

//Эндпоинт /addOrRefactor
func (r *Rest) addOrRefactor(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start refactoring Product",
	}).Print()

	final := models.ProductDataInMS{}
	name := strings.Replace(c.FormValue("name"), "@", " ", -1) //Перевести то потом в структуру ,Сделать валидацию
	if name == "" {
		r.Logger.Error(fmt.Errorf("поле name пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	nomecl := c.FormValue("nomencl")
	if nomecl == "" {
		r.Logger.Error(fmt.Errorf("поле namecl пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	count, err := strconv.Atoi(c.FormValue("count"))
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if count == 0 {
		r.Logger.Error(fmt.Errorf("поле namecl пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	idSupply := c.FormValue("id")
	if nomecl == "" {
		r.Logger.Error(fmt.Errorf("поле id пустует"))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	whatDo := c.FormValue("whatDo")
	if whatDo == "new" {
		final, err = r.MakeProduct(nomecl, price, name, c)
		if err != nil {
			r.Logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else if whatDo == "refactor" {
		final, err = r.RefactorProduct(nomecl, name, c)
		if err != nil {
			r.Logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	final.Count = count
	err = addPositionInSupply(final, idSupply, r.Token.Access_token) //Добавление продукта в приемку
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return c.JSON(http.StatusOK, "ok")
}

// Относящиеся к createFile ф-ции методы

func createFileInput(file multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(proto.NEW_FILE_XLS)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
func createFileCsv() (string, error) {
	new, err := addNewConvertFile()
	if err != nil {
		return "", err
	}

	final, err := finalNewConvertFile(new.Data.Id)
	if err != nil {
		return "", err
	}

	time.Sleep(3 * time.Second)

	url, err := takeUrl(final.Data.Id)
	if err != nil {
		return "", err
	}

	err = downloadFile(proto.NEW_FILE_CSV, url.Data.Output.Url)
	if err != nil {
		return "", err
	}
	return proto.NEW_FILE_CSV, err
}

//
func addNewConvertFile() (*models.ConvertFileResponse, error) {
	var result = new(models.ConvertFileResponse)
	data := models.ConvertFile{
		Apikey:       "a86cdbd799a71c05f50b8a2bc556515e",
		Input:        "upload",
		OutputFormat: "CSV",
	}

	mes, err := json.Marshal(data) //Кодирование
	if err != nil {
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.convertio.co/convert", bytes.NewBuffer(mes))
	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var er = new(models.ConvertFileError)
	err = json.Unmarshal([]byte(bodyText), &result)
	if err != nil {
		return result, err
	}

	if result.Status != "ok" {
		return result, fmt.Errorf("%v", er.Err)
	}
	return result, err
}
func finalNewConvertFile(id string) (*models.ConvertFileFinal, error) {
	result := new(models.ConvertFileFinal)

	data, err := os.Open(proto.NEW_FILE_XLS)
	if err != nil {
		return result, err
	}

	client := &http.Client{}
	url := "http://api.convertio.co/convert/" + id + "/" + "Товары.xls"
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var er = new(models.ConvertFileError)
	json.Unmarshal([]byte(bodyText), &result)
	if result.Status != "ok" {
		json.Unmarshal([]byte(bodyText), &er)
		return result, fmt.Errorf(er.Err)
	}
	return result, err
}
func takeUrl(id string) (*models.ReadyFile, error) {
	result := new(models.ReadyFile)

	client := &http.Client{}
	url := "https://api.convertio.co/convert/" + id + "/" + "status"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var er = new(models.ConvertFileError)
	json.Unmarshal([]byte(bodyText), &result)
	if result.Status != "ok" {
		json.Unmarshal([]byte(bodyText), &er)
		return result, fmt.Errorf(er.Err)
	}
	return result, err
}
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//
