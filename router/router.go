package router

import (
    "encoding/json"
    "net/http"

    "github.com/dllg/go-funny-endpoints/funny"
    "github.com/dllg/go-funny-endpoints/httpclient"
)

type msgfunc func(httpclient.HTTPClient) string

func messageHandler(f msgfunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        msg := f(&hc)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": msg})
    }
}

var (
    hc httpclient.Impl
)

// Setup will setup all endpoints handling different http requests
func Setup() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/v1/advice", messageHandler(funny.GetAdviceFromAPI))
    mux.HandleFunc("/v1/chucknorris", messageHandler(funny.GetChuckNorrisJokeFromAPI))
    mux.HandleFunc("/v1/dadjoke", messageHandler(funny.GetDadJokeFromAPI))
    mux.HandleFunc("/v1/random", messageHandler(funny.GetRandomMessage))
    return mux
}
