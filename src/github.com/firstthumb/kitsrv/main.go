package main

import(
  "net/http"
  "os"
	"os/signal"
  "syscall"

  "fmt"

  "github.com/firstthumb/kitsrv/github"

  "github.com/go-kit/kit/log"
)

func main() {
  var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

  var svc github.GithubService
  {
    svc = github.New()
  }

  var endpoints github.Endpoints
  {
    endpoints = github.Endpoints { GithubEndpoint: github.GithubEndpoint(svc) }
  }

  errc := make(chan error)

  // Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

  // HTTP transport.
	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", ":8000")

		handler := github.MakeHttpHandler(endpoints, logger)
		errc <- http.ListenAndServe(":8000", handler)
	}()

  logger.Log("exit", <-errc)
}
