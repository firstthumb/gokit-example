package service

import (
	//"fmt"
	"context"
	"encoding/json"
	"net/http"
  "github.com/firstthumb/usersrv/requests"
  "github.com/firstthumb/usersrv/responses"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(eps *Endpoints) http.Handler {
	h := mux.NewRouter()
	op := httptransport.ServerErrorEncoder(encodeError)

	h.Methods("GET").Path("/").Handler(httptransport.NewServer(
		eps.Get,
		decodeGet,
		encodeResp,
		op))
	h.Methods("POST").Path("/").Handler(httptransport.NewServer(
		eps.Add,
		decodeAdd,
		encodeResp,
		op))
	h.Methods("PUT").Path("/").Handler(httptransport.NewServer(
		eps.Update,
		decodeUpdate,
		encodeResp,
		op))
	h.Methods("DELETE").Path("/").Handler(httptransport.NewServer(
		eps.Delete,
		decodeDelete,
		encodeResp,
		op))

	return h
}

func decodeGet(_ context.Context, r *http.Request) (interface{}, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	id := r.Form.Get("user_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return requests.Get{id_int}, nil
}

func decodeAdd(_ context.Context, r *http.Request) (interface{}, error) {
	req := requests.Add{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	id := r.Form.Get("user_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	req := requests.Update{}
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	req.ID = id_int
	return req, nil
}

func decodeDelete(_ context.Context, r *http.Request) (interface{}, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	id := r.Form.Get("user_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return requests.Delete{id_int}, nil
}

func encodeResp(_ context.Context, w http.ResponseWriter, response interface{}) error {
	g := responses.General{response, ""}
	return json.NewEncoder(w).Encode(g)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	g := responses.General{nil, err.Error()}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(g)
}
