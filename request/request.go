//Выполнение запроса к dadata.
package request

import (
	"net/http"
	"net/url"
)

const GET = "GET"
const POST = "POST"

type Request struct {
	Url         string
	Method      string
	QueryParams url.Values
	PostData    map[string]interface{}
	client      *http.Client
	request     *http.Request
}

func (r *Request) exec() error {
	return nil
}
