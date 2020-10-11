package dadata

import (
	"os"
	"testing"
	"time"
)

func config() Config {
	return Config{
		Token:   os.Getenv("DADATA_KEY"),
		Timeout: time.Duration(10) * time.Second,
	}
}

func TestDadata_Suggestions(t *testing.T) {
	var config = config()
	New(&config).Suggestions()
}
