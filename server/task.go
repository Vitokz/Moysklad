package server

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Vitokz/Moysklad/models"
	"github.com/Vitokz/Moysklad/proto"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

//Эндпоинт /
func (r *Rest) GetTask(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Test endpoint",
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
	token := models.Token{}
	json.Unmarshal([]byte(bodyText), &token) //Перезаписываю []byte в JSON

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

	/*resp*/
	_, err = client.Do(req)
	if err != nil {
		r.Logger.Errorf("Failed processing take products request: %s\n", err)
		return err
	}
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	r.Logger.WithError(err).Error()
	// 	return err
	// }
	return nil
}

//Эндпоинт /createPrice
func (r *Rest) CreatePrice(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start creationg pricelist",
	})
	file, err := openCSV("Предварительный_счёт.csv")
	if err != nil {
		r.Logger.Error(err)
	}
	productsInMS, err := r.CSVtakeProducts()
	if err != nil {
		r.Logger.Error(err)
	}
	name, err := createFinalCsv(file, *productsInMS)
	if err != nil {
		r.Logger.Error(err)
	}
	r.Logger.Println()
	return c.JSON(http.StatusOK, name)
}

func openCSV(name string) (*os.File, error) {
	file, err := os.Open("proto/priceList/" + name)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (r *Rest) CSVtakeProducts() (*models.CSVRows, error) {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start to taking products MS (CSV)",
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
	result := new(models.CSVRows)
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	json.Unmarshal([]byte(bodyText), &result)
	return result, nil
}
func createFinalCsv(file *os.File, products models.CSVRows) (string, error) {
	new, err := os.Create("result.csv")
	if err != nil {
		return "", err
	}
	writer := csv.NewWriter(new)
	defer writer.Flush()
	allWriters, err := checkSimilar(file, products)
	if err != nil {
		return "", err
	}
	file.Close()
	gocsv.MarshalFile(allWriters, new)
	return "result.csv", nil
}
func checkSimilar(file *os.File, products models.CSVRows) (*[]models.CSVProductsFinal, error) {
	//Для предварительного счета поля: 0-номер,1-имя,6-цена,8-кол-во,9-вид ед измерения,10-сумма
	result := make([]models.CSVProductsFinal, 0)

	reader := csv.NewReader(file)
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		tp, _ := strconv.Atoi(record[0])
		if tp == 0 {
			continue
		} else {
		loop:
			for _, v := range products.Rows {
				aliases := strings.Split(v.Description, ";")
				for _, s := range aliases {
					if s == record[1] {
						count, err := strconv.ParseFloat(record[8], 64)
						if err != nil {
							return nil, err
						}
						price, err := strconv.ParseFloat(record[6], 64)
						if err != nil {
							return nil, err
						}
						result = append(result, models.CSVProductsFinal{
							Number: record[0],
							Name:   v.Name,
							Count:  int(count),
							Price:  int(price),
							Summ:   int(count * price),
						})
						break loop
					}
				}
			}
		}
	}
	return &result, nil
}

//Эндпоинт /makeSupply
func (r *Rest) MakeSupply(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start creating Supply",
	})
	supply, err := makeNewSupplyInMS(r.Token.Access_token) //Создание новой приемки в мойсклад
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}
	file, err := openCSV("Предварительный_счёт.csv") //Открытие csv файла поставщика
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}
	products, err := CSVRows(file) //Парсинг файла(Достает все строки с товарами и берет нужную информацию)
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}
	err = findProductsInMs(*products, r.Token.Access_token, supply.Id) //Поиск продукта в мойсклад и в случае успеха добавление в приемку
	if err != nil {
		r.Logger.Errorf("Failed: %v", err)
		return err
	}
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	})
	return c.JSON(http.StatusOK, supply)
}

func makeNewSupplyInMS(id string) (models.Supply, error) {
	supplyData := models.MakeNewSupply() //Создание модели с данными приемки (на данный момент все заполнено статически)
	client := &http.Client{}             //Создание клиента запроса

	mes, err := json.Marshal(supplyData) //Первод в []byte модели приемки
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
	result := new(models.Supply) // Модель Supply Содержит в себе только id приемки
	json.Unmarshal([]byte(bodyText), result)
	return *result, nil
} //+
func CSVRows(file *os.File) (*[]models.CsvProducts, error) {
	defer file.Close() //Закрытие файла по окончании ф-ции
	//Для предварительного счета поля: 0-номер,1-имя,6-цена,8-кол-во,9-вид ед измерения,10-сумма
	result := make([]models.CsvProducts, 0) //Создание массива моделей товаров поставщика csv

	reader := csv.NewReader(file) //Чтение файла
	for {
		record, e := reader.Read() //Беру одну строку
		if e != nil {
			fmt.Println(e)
			break
		}
		tp, _ := strconv.Atoi(record[0]) //Проверяю первый символ строки на наличие числа ,так как это отлич знак товара
		if tp == 0 {
			continue
		} else {
			count, err := strconv.ParseFloat(record[8], 64)
			if err != nil {
				return nil, err
			}
			result = append(result, models.CsvProducts{
				Count: int(count),
				Name:  record[1],
			})
		}
	}
	return &result, nil
} //+
func findProductsInMs(products []models.CsvProducts, idUser string, idSupply string) error {
	for _, v := range products {
		check, err, ok := searchProduct(v, idUser) //Поиск продукта в списке товаров в Мойсклад
		if err != nil {
			return err
		}
		if ok {
			err := addPositionInSupply(check, idSupply, idUser) //Добавление продукта в приемку
			if err != nil {
				return err
			}
		}
	}
	return nil
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
	_ = bodyText
	fmt.Println(string(bodyText))
	return nil
}

// Эндпоинт /makeProduct
func (r *Rest) MakeProduct(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start creating Product",
	}).Print()
	product := models.NewProductModel("Darova", 64)
	err := addProductInMS(product, r.Token.Access_token)
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return c.JSON(http.StatusOK, "all okey")
}

func addProductInMS(product *models.NewProduct, id string) error {

	mes, err := json.Marshal(product) //Кодирование
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", proto.PRODUCTS_IN_MY_SKLAD, bytes.NewBuffer(mes))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+id)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_ = bodyText
	return nil
}

// Эндпоинт /refactorProduct
func (r *Rest) RefactorProduct(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Start refactoring Product",
	}).Print()
	product, err := searchProductForRefactor("Drova", r.Token.Access_token)
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	product.Description += ";Darova esd"

	err = refactorProduct(product, r.Token.Access_token)
	if err != nil {
		r.Logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	r.Logger.WithFields(logrus.Fields{
		"status": "ok",
	}).Println()
	return c.JSON(http.StatusOK, "all okey")
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

	rows := new(models.RowsForRefactor) //Создание модели для получения ответа от мойсклад
	json.Unmarshal([]byte(bodyText), &rows)
	if len(rows.Rows) == 0 {
		return models.RefactorProductsInMS{}, fmt.Errorf("no find products")
	}
	result := rows.Rows[0]
	return result, nil //Возврат готовой модели товара из МойСклад
}

func refactorProduct(product models.RefactorProductsInMS, id string) error {
	client := &http.Client{} //Создание клиента запроса

	mes, err := json.Marshal(product) //Первод в []byte модели товара
	if err != nil {
		return err
	}

	url := proto.PRODUCTS_IN_MY_SKLAD + "/" + product.Id          //Создание урла запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(mes)) //Сам запрос
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+id)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	/*resp*/
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	r.Logger.WithError(err).Error()
	// 	return err
	// }
	return nil
}
