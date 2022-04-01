package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"weixin-golang/pkg"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getIndex()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprint(w, data)
}

func StockHandler(w http.ResponseWriter, r *http.Request) {
	result, err := pkg.FetchLastPrice("sz399001", "sz002594", "sz002230")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", result)))
	}
}

func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
