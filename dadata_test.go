package dadata

import (
	"os"
	"testing"
)

func config() Config {
	os.Setenv("KEY", "061863cf0ac433f1e3e918f5d7f88d50a3b3af24")
	return Config{
		Token:   os.Getenv("KEY"),
		Timeout: 10,
	}
}

func TestDadata_Suggestions(t *testing.T) {
	var config = config()
	New(&config).Suggestions()
}

func TestDadata_Suggestions_Address(t *testing.T) {
	var config = config()
	var d = New(&config)
	items, err := d.Suggestions().Address("Москва Мытная 7", 10)
	if err != nil {
		t.Error(err)
	}
	if len(items) != 8 {
		t.Errorf(`Waiting for %d results, got %d`, 8, len(items))
	}
}
