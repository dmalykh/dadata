package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Token  string
	Post   map[string]interface{}
	Client *http.Client
	Handle func(ctx context.Context, request Request, v *interface{}) error
}

//Метод выполняет запрос в dadata и делает unmarshal результата в v, используя специальный handler для выполнения запроса.
func (c Client) Request(ctx context.Context, req Request, v interface{}) error {
	u, err := url.Parse(req.Url)
	if err != nil {
		return fmt.Errorf(`Can't parse url "%s": %s`, req.Url, err.Error())
	}
	postData, err := json.Marshal(req.PostData)
	if err != nil {
		return fmt.Errorf(`Can't marshal post "%#v": %s`, req.PostData, err.Error())
	}

	request, err := http.NewRequest(req.Method, u.String(), bytes.NewReader(postData))
	if err != nil {
		return fmt.Errorf(`Can't create new request "%s": %s`, u.String(), err.Error())
	}

	request.WithContext(ctx)
	request.Header.Set("Authorization", "Token "+c.Token)
	request.Header.Set("Content-Type", "application/json")

	req.request = request
	req.client = c.Client
	return c.Handle(ctx, req, &v)
}
