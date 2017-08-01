package handlers

import (
	"net/http"
	"strconv"
)

func GetString(request *http.Request, name string) string {
	return request.URL.Query().Get(name)
}

func GetInteger(request *http.Request, name string) (value int64) {
	value, _ = strconv.ParseInt(GetString(request, name), 10, 64)
	return
}

func GetStringArray(request *http.Request, name string) []string {
	return request.URL.Query()[name]
}