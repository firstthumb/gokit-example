package main

import (
	"net/http"
	"os"
	"github.com/firstthumb/usersrv/service"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)
	svc := service.New()
	eps := service.MakeEndpoints(svc, logger)
	h := service.MakeHTTPHandler(eps)

	http.ListenAndServe(":8080", h)
}
