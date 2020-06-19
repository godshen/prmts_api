package controller

import (
	"encoding/json"
	"net/http"
)

// ResponseHandler is error handler for http
type ResponseHandler func(w http.ResponseWriter, r *http.Request) (bool, interface{})

func (h ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		OK   bool        `json:"ok"`
		Data interface{} `json:"data"`
	}
	ok, ret := h(w, r)
	res := Response{ok, ret}
	byteRes, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(byteRes)
}
