//Методы для работы с API Dadata.
//Тип Config содержит метод Handle, который можно переопределить для изменения выполнения самих запросов в API dadata.ru, например, для кэширования.
package dadata

import (
	"context"
	"github.com/dmalykh/dadata/request"
	"github.com/dmalykh/dadata/suggestions"
	"net/http"
	"time"
)

type Config struct {
	Token   string                                                                   //Токен для обращения к API
	Timeout time.Duration                                                            //Время для timeout запроса в dadata
	Handle  func(ctx context.Context, request request.Request, v *interface{}) error //Метод, через который будет "проходить" ответ с сервиса
}

type Dadata struct {
	request *request.Client
}

//Возвращает новый экземпляр dadata
func New(config *Config) *Dadata {
	if config.Handle == nil {
		config.Handle = request.DefaultHandler
	}

	return &Dadata{
		request: &request.Client{
			Token:  config.Token,
			Handle: config.Handle,
			Client: &http.Client{
				Timeout: config.Timeout,
				Transport: &http.Transport{
					MaxIdleConnsPerHost: 1024,
					TLSHandshakeTimeout: 0 * time.Second,
				},
			},
		},
	}
}

//Возвращает экземплятр структуры через singletone для работы с подсказками
func (d *Dadata) Suggestions() *suggestions.Suggestions {
	return suggestions.GetInstance(d.request)
}
