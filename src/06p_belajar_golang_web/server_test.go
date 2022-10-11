package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:9090",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
