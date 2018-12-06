package suggestions

import (
	"dadata/request"
	"fmt"
)

type Suggestions struct {
	Config *request.Config
}

func GetInstance(c *request.Config) *Suggestions {
	return &Suggestions{
		Config: c,
	}
}

const DADATA_SUGGESTIONS_URL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/%s"

//Метод для выполнения запросов типа suggestions в dadata
func (s *Suggestions) request(kind string, post map[string]interface{}, result interface{}) error {
	var u = fmt.Sprintf(DADATA_SUGGESTIONS_URL, kind)
	if err := s.Config.Request(u, post, result); err != nil {
		return fmt.Errorf(`Can't make request to "%s": %s`, u, err.Error())
	}
	return nil
}
