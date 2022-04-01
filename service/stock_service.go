package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weixin-golang/pkg"
)

type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

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
	res := &JsonResult{}
	result, err := pkg.FetchLastPrice("sz399001", "sz002594", "sz002230")
	var b []byte
	if err == nil {
		res.Code = 200
		res.Data = fmt.Sprintf("%v", result)
		b, err = json.Marshal(res)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
