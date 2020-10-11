package request

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"testing"
	"time"
)

func TestKeepAlive(t *testing.T) {
	var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/keep" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}
	}))

	var newConnectionsCounter = 0
	var traceContext = httptrace.WithClientTrace(context.TODO(), &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				newConnectionsCounter++
			}
		},
	})

	var client = getClient()

	client.Request(traceContext, Request{
		Url:    server.URL + "/keep",
		Method: GET,
	}, nil)
	client.Request(traceContext, Request{
		Url:    server.URL + "/keep",
		Method: POST,
	}, nil)
	client.Request(traceContext, Request{
		Url:    server.URL + "/keep",
		Method: GET,
	}, nil)

	if newConnectionsCounter != 1 {
		t.Errorf(`Waitting for 1 connection, got %d`, newConnectionsCounter)
	}
}

func BenchmarkRequest(b *testing.B) {
	var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/keep" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}
	}))
	var client = getClient()

	b.Run("client", func(b *testing.B) {
		b.ReportAllocs()
		client.Request(context.TODO(), Request{
			Url:    server.URL + "/keep",
			Method: GET,
		}, nil)
	})
}

func getClient() Client {
	return Client{
		Handle: func(ctx context.Context, request Request, v *interface{}) error {
			return DefaultHandler(ctx, request, v)
		},
		Client: &http.Client{
			Timeout: time.Duration(3600) * time.Second, //Для дебаггера
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 1024,
				TLSHandshakeTimeout: 0 * time.Second,
			},
		},
	}
}
