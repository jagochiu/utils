package webtools

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
SentRequest -
*/
func SentRequest(url string, method string, data DataObj, headers map[string]string) (string, int) {
	var err error
	var buf bytes.Buffer
	var req *http.Request
	requestData := &bytes.Buffer{}
	if method != "GET" {
		switch data.Zip {
		case `gzip`:
			headers["Content-Encoding"] = "gzip"
			if g, err := gzip.NewWriterLevel(&buf, 9); err != nil {
			} else if _, err := g.Write([]byte(data.Data)); err != nil {
				fmt.Println(err)
			} else if err := g.Close(); err != nil {
				fmt.Println(err)
			}
			requestData = &buf
		default:
			requestData = bytes.NewBuffer([]byte(data.Data))
		}
		method = `POST`
	}
	req, err = http.NewRequest(method, url, requestData)
	if err != nil {
		log.Printf("%v \n ", err)
	}
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	timeout := time.Duration(15 * time.Second)
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("URL:> ", url)
		fmt.Println(resp)
		fmt.Println(err)
		return "ERROR", -1
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("URL:> ", url)
		fmt.Println("Response Status:", resp.Status)
		return "", resp.StatusCode
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "ERROR", -1
	}
	return string(body), 1
}

/*
HeaderType -
*/
type HeaderType struct {
	Key   string
	Value string
}

/*
DataType -
*/
type DataObj struct {
	Zip  string
	Data string
}
