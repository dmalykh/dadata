package request

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DadataError struct {
	Family  string `json:"family"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

//Метод по-умолчанию, выполняющий запросы в дадата.
func DefaultHandler(ctx context.Context, request Request, v *interface{}) error {

	resp, err := request.client.Do(request.request)
	if err != nil {
		return fmt.Errorf(`Can't make request to "%s": %s`, request.Url, err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(`Can't read response from "%s": %s`, request.Url, err.Error())
	}
	//Проверка не вернула ли dadata ошибку
	var dadataerr DadataError
	if json.Unmarshal(body, &dadataerr); dadataerr.Reason != "" {
		return fmt.Errorf(`Error from dadata: %s`, string(body))
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf(`Can't unmarshal result "%s" from "%s" with : %s`, string(body), request.Url, err.Error())
	}
	return nil
}
