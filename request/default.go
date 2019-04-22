package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DadataError struct {
	Family  string `json:"family"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

//Метод по-умолчанию, выполняющий запросы в дадата.
func DefaultHandler(c DadataRequest, w *interface{}) error {
	u, err := url.Parse(c.ApiUrl)
	if err != nil {
		return fmt.Errorf(`Can't parse url "%s": %s`, c.ApiUrl, err.Error())
	}
	data, err := json.Marshal(c.Post)
	if err != nil {
		return fmt.Errorf(`Can't marshal post "%#v": %s`, c.Post, err.Error())
	}

	var client = http.Client{
		Timeout: c.Timeout,
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf(`Can't create new request "%s": %s`, u.String(), err.Error())
	}

	req.WithContext(c.Ctx)
	req.Header.Set("Authorization", "Token "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf(`Can't make request to "%s" with "%s": %s`, u.String(), string(data), err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(`Can't read response from "%s" with "%s": %s`, u.String(), string(data), err.Error())
	}
	//Проверка не вернула ли dadata ошибку
	var dadataerr DadataError
	if json.Unmarshal(body, &dadataerr); dadataerr.Reason != "" {
		return fmt.Errorf(`Error from dadata: %s`, string(body))
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		return fmt.Errorf(`Can't unmarshal result "%s" from "%s" with "%s": %s`, string(body), u.String(), string(data), err.Error())
	}
	return nil
}
