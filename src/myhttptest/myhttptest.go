package myhttptest

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", SendJson)
}

func SendJson(rw http.ResponseWriter, r *http.Request) {

	u := struct {
		Name string
	}{
		Name: "张三",
	}

	rw.Header().Set("Content-Type", "application-json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)

}
