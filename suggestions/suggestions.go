package suggestions

import (
	"context"
	"fmt"
	"github.com/dmalykh/dadata/request"
)

type Suggestions struct {
	Config *request.DadataRequest
}

func GetInstance(c *request.DadataRequest) *Suggestions {
	return &Suggestions{
		Config: c,
	}
}

var DADATA_SUGGESTIONS_URL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/%s"

//Метод для выполнения запросов типа suggestions в dadata
func (s *Suggestions) request(ctx context.Context, kind string, post map[string]interface{}, result interface{}) error {
	var u = fmt.Sprintf(DADATA_SUGGESTIONS_URL, kind)
	if err := s.Config.Request(ctx, u, post, result); err != nil {
		return fmt.Errorf(`Can't make request to "%s": %s`, u, err.Error())
	}
	return nil
}
