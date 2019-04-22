//Выполнение запроса к dadata.
package request

import (
	"time"
)

type DadataRequest struct {
	Token   string
	Timeout time.Duration
	ApiUrl  string
	Post    map[string]interface{}
	Handle  func(c DadataRequest, w *interface{}) error
}

//Метод выполняет запрос в dadata и делает unmarshal результата в v, используя специальный hander для выполнения запроса.
func (r DadataRequest) Request(apiUrl string, post map[string]interface{}, v interface{}) error {
	r.ApiUrl = apiUrl
	r.Post = post
	return r.Handle(r, &v)
}
