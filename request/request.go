//Выполнение запроса к dadata.
package request

import (
	"context"
	"time"
)

type DadataRequest struct {
	Token   string
	Timeout time.Duration
	ApiUrl  string
	Post    map[string]interface{}
	Handle  func(c DadataRequest, w *interface{}) error
	Ctx     context.Context
}

//Метод выполняет запрос в dadata и делает unmarshal результата в v, используя специальный hander для выполнения запроса.
func (r DadataRequest) Request(ctx context.Context, apiUrl string, post map[string]interface{}, v interface{}) error {
	r.ApiUrl = apiUrl
	r.Post = post
	r.Ctx = ctx
	return r.Handle(r, &v)
}
