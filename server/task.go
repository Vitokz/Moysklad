package server

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Vitokz/Moysklad/models"
	"github.com/Vitokz/Moysklad/proto"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (r *Rest) GetTask(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "hello world",
	}).Println()
	// ctx := c.Request().Context()

	// err := r.Handler.GetTask(ctx, useruuid)
	// if err != nil {

	// }
	// return c.JSON()
	return c.JSON(http.StatusOK, "Good")
}

func (r *Rest) Auth(c echo.Context) error {
	r.Logger.WithFields(logrus.Fields{
		"event": "Authorization",
	}).Println()

	client := &http.Client{}

	req, err := http.NewRequest("POST", proto.BASIC_URL_MOYSKLAD, nil)
	if err != nil {
		r.Logger.WithError(err).Error()
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(proto.LOGIN, proto.PASSWORD))
	resp, err := client.Do(req)
	if err != nil {
		r.Logger.Errorf("Failed processing add login request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	token := models.Token{}
	json.Unmarshal([]byte(bodyText), &token)

	r.Logger.WithFields(logrus.Fields{
		"statusAuth": "OK",
	}).Println()

	return c.JSON(http.StatusOK, token)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}
