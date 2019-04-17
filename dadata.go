//Методы для работы с API Dadata.
//Тип Config содержит метод Handle, который можно переопределить для изменения выполнения самих запросов в API dadata.ru, например, для кэширования.
package dadata

import (
	"github.com/dmalykh/dadata/request"
	"github.com/dmalykh/dadata/suggestions"
	"sync"
	"time"
)

type Config struct {
	Token   string                                             //Токен для обращения к API
	Timeout time.Duration                                      //Время для timeout запроса в dadata
	Handle  func(c request.DadataRequest, w interface{}) error //Метод, через который будет "проходить" ответ с сервиса
}

type Dadata struct {
	request    *request.DadataRequest
	suggestion struct {
		once sync.Once
		s    *suggestions.Suggestions
	}
}

//Возвращает новый экземпляр dadata
func New(config *Config) *Dadata {
	if config.Handle == nil {
		config.Handle = request.DefaultHandler
	}
	return &Dadata{
		request: &request.DadataRequest{
			Token:   config.Token,
			Timeout: time.Duration(config.Timeout) * time.Second,
			Handle:  config.Handle,
		},
	}
}

//Возвращает экземплятр структуры через singletone для работы с подсказками
func (d *Dadata) Suggestions() *suggestions.Suggestions {
	d.suggestion.once.Do(func() {
		d.suggestion.s = suggestions.GetInstance(d.request)
	})
	return d.suggestion.s
}
