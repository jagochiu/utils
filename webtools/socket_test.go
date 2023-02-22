package webtools

import (
	"errors"
	"io/ioutil"
	"testing"
)

func TestSocket(t *testing.T) {
	dialer := Socks5Client(TCP, `127.0.0.1`, 9001)
	if dialer == nil {
		panic(errors.New(`socket fail`))
	}
	resp, err := dialer.Get("https://wtfismyip.com/json")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(errors.New(`body read failed`))
	}
	println(string(body))
}
func TestRequestSocket(t *testing.T) {
	resp, status := SentRequest(`https://wtfismyip.com/json`, `GET`, DataObj{
		Socket: Socket{
			Used:     true,
			Protocol: TCP,
			Host:     `127.0.0.1`,
			Port:     9101,
		},
	}, map[string]string{})
	if status == 1 {
		println(resp)
	}
}
