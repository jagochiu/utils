package webtools

import (
	"testing"
)

func TestRequestGET(t *testing.T) {
	resp, status := SentRequest(`https://ipinfo.io/json`, `GET`, DataObj{}, map[string]string{})
	if status == 1 {
		println(resp)
	}
}
func TestRequestPOST(t *testing.T) {
}
