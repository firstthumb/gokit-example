package service

import(
  "context"

  "github.com/firstthumb/usersrv/requests"

  "github.com/go-kit/kit/log"
  "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
  Get endpoint.Endpoint
  Add endpoint.Endpoint
  Update endpoint.Endpoint
  Delete endpoint.Endpoint
}

func MakeEndpoints(svc UserService, logger log.Logger) *Endpoints {
  eps := &Endpoints{}
	eps.Add = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Add)
		return svc.Add(&r)
	}
	eps.Get = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Get)
		return svc.Get(&r)
	}
  eps.Update = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Update)
		return svc.Update(&r)
	}
	eps.Delete = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Delete)
		return svc.Delete(&r)
  }

  eps.Get = LoggingMiddleware(logger)(eps.Get)

	return eps
}
