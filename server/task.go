package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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

func basicAuth(username, password string) string { // Функция относящаяся к auth
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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
