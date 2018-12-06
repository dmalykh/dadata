//Методы для работы с API Dadata
package dadata

import (
	"dadata/request"
	"dadata/suggestions"
	"sync"
)

type Dadata struct {
	Config     *request.Config
	suggestion struct {
		once sync.Once
		s    *suggestions.Suggestions
	}
}

func New(config *request.Config) *Dadata {
	return &Dadata{
		Config: config,
	}
}

//Возвращает экземплятр структуры через singletone для работы с подсказками
func (d *Dadata) Suggestions() *suggestions.Suggestions {
	d.suggestion.once.Do(func() {
		d.suggestion.s = suggestions.GetInstance(d.Config)
	})
	return d.suggestion.s
}
