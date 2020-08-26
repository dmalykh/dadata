package suggestions

import (
	"context"
	"fmt"
	"github.com/dmalykh/dadata/request"
)

type Suggestions struct {
	Client *request.Client
}

func GetInstance(c *request.Client) *Suggestions {
	return &Suggestions{
		Client: c,
	}
}

var DADATA_SUGGESTIONS_URL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/%s"

//Метод для выполнения запросов типа suggestions в dadata
func (s *Suggestions) makeRequest(ctx context.Context, method string, req request.Request, result interface{}) error {
	req.Url = fmt.Sprintf(DADATA_SUGGESTIONS_URL, method)
	if err := s.Client.Request(ctx, req, result); err != nil {
		return fmt.Errorf(`Can't make makeRequest to "%s": %s`, req.Url, err.Error())
	}
	return nil
}
