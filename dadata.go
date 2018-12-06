//Методы для работы с API Dadata
package dadata

import (
	"github.com/dmalykh/dadata/request"
	"github.com/dmalykh/dadata/suggestions"
	"sync"
	"time"
)

type Config struct {
	Token   string //Токен для обращения к API
	Timeout uint
}

type Dadata struct {
	request    request.Request
	suggestion struct {
		once sync.Once
		s    *suggestions.Suggestions
	}
}

//Возвращает новый экземпляр dadata
func New(config *Config) *Dadata {
	return &Dadata{
		request: request.Request{
			Token:   config.Token,
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

//Возвращает экземплятр структуры через singletone для работы с подсказками
func (d *Dadata) Suggestions() *suggestions.Suggestions {
	d.suggestion.once.Do(func() {
		d.suggestion.s = suggestions.GetInstance(&d.request)
	})
	return d.suggestion.s
}
