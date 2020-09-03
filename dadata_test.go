package dadata

import (
	"context"
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

func TestDadata_Suggestions_Address(t *testing.T) {
	var config = config()
	var d = New(&config)
	var ctx = context.Background()
	items, err := d.Suggestions().Address(ctx, "Москва Мытная 7", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 8 {
		t.Errorf(`Waiting for %d results, got %d`, 8, len(items))
	}
}
