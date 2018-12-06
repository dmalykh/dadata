package dadata

import (
	"dadata/request"
	"testing"
)

func TestDadata_Suggestion(t *testing.T) {
	var config = request.Config{
		Token:   "",
		Timeout: 10,
	}
	New(&config).Suggestions()
}
