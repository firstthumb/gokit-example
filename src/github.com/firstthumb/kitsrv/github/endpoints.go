package github

import(
  "context"

  "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GithubEndpoint     endpoint.Endpoint
}

func GithubEndpoint(svc GithubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(githubSearchRequest)
		v, err := svc.GetRepositories(ctx, req.S)
		if err != nil {
			return githubSearchResponse{v, err.Error()}, nil
		}
		return githubSearchResponse{v, ""}, nil
	}
}


type githubSearchRequest struct {
  S     string      `json:"s"`
}

type githubSearchResponse struct {
  V     string      `json:"v"`
  Error string      `json:"err,omitempty"`
}
