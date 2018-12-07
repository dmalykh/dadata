package dadata

import (
	"testing"
)

func TestDadata_Suggestion(t *testing.T) {
	var config = Config{
		Token:   "",
		Timeout: 10,
	}
	New(&config).Suggestions()
}
