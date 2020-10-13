package 模拟调用

import (
	"log"
	"myhttptest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	myhttptest.Routes()
}

func TestSendJson(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)

	if err != nil {
		t.Fatal("创建Request失败")
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("code:", rw.Code)

	log.Println("body:", rw.Body.String())

}
