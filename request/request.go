package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	Token   string //Токен для обращения к API
	Timeout uint
}

//Метод выполняет запрос в dadata и делает unmarshal результата в v
func (c *Config) Request(apiUrl string, post map[string]interface{}, v interface{}) error {

	u, err := url.Parse(apiUrl)
	if err != nil {
		return fmt.Errorf(`Can't parse url "%s": %s`, apiUrl, err.Error())
	}
	data, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf(`Can't marshal post "%#v": %s`, post, err.Error())
	}

	var client = http.Client{
		Timeout: time.Duration(c.Timeout) * time.Second,
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf(`Can't create new request "%s": %s`, u.String(), err.Error())
	}
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
	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf(`Can't unmarshal result "%s" from "%s" with "%s": %s`, string(body), u.String(), string(data), err.Error())
	}
	return nil
}
