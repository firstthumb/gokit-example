package github

import(
  "context"
  "encoding/json"
  "net/http"

  "errors"

  "github.com/go-kit/kit/log"

  httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(endpoints Endpoints, logger log.Logger) http.Handler {
  options := []httptransport.ServerOption{
    httptransport.ServerErrorLogger(logger),
  }
  m := http.NewServeMux()
  m.Handle("/github",
    httptransport.NewServer(
      endpoints.GithubEndpoint,
      DecodeHttpGithubSearchRequest,
      EncodeHttpResponse,
      options...,
  ))

  return m
}

func DecodeHttpGithubSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request githubSearchRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    return nil, err
  }
  return request, nil
}

func EncodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

func errorDecoder(r *http.Response) error {
  var w errorWrapper
  if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
    return err
  }
  return errors.New(w.Error)
}

type errorWrapper struct {
  Error string `json:"error"`
}
