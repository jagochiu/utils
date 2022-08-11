package utils

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jagochiu/utils/jrsa"
)

/*
Received -
*/
type Received struct {
	ID          string      `json:"_id,omitempty" bson:"_id,omitempty"`
	Merchant    string      `json:"merchant,omitempty" bson:"merchant,omitempty"`
	MerchantKey string      `json:"key,omitempty" bson:"key,omitempty"`
	Action      string      `json:"action,omitempty" bson:"action,omitempty"`
	Auth        string      `json:"auth,omitempty" bson:"auth,omitempty"`
	Date        int64       `json:"date,omitempty" bson:"date,omitempty"`
	PayLoad     interface{} `json:"payload,omitempty" bson:"payload,omitempty"`
}

/*
ToObject -
*/
func (r *Received) ToObject(obj interface{}) bool {
	byteJSON, err := json.Marshal(r.PayLoad)
	if err != nil {
		log.Printf("ReceivedToObject Marshal fail %v \n", err)
		return false
	}
	// println(string(byteJSON))
	err = json.Unmarshal(byteJSON, &obj)
	if err != nil {
		log.Printf("ReceivedToObject Unmarshal fail %v \n", err)
		return false
	}
	return true
}

/*
ToObject -
*/
func (r *Received) Decrypte(c *gin.Context, pri string) error {
	body, err := GetBody(c)
	if err != nil {
		log.Printf("%v \n", err)
		return err
	}
	if len(body) <= 0 {
		return errors.New(`empty body`)
	}
	rs, err := jrsa.PriDecrypt(string(body), pri)
	if err != nil {
		return err
	}
	if len(rs) <= 0 {
		return errors.New(`empty decrypted result`)
	}
	byteJSON, err := json.Marshal(rs)
	if err != nil {
		log.Printf("ReceivedToObject Marshal fail %v \n", err)
		return err
	}
	// println(string(byteJSON))
	err = json.Unmarshal(byteJSON, &r)
	if err != nil {
		log.Printf("ReceivedToObject Unmarshal fail %v \n", err)
		return err
	}
	return nil
}

/*
GetBody -
*/
func GetBody(c *gin.Context) (bytesBody []byte, err error) {
	bytesBody = []byte(``)
	bytesBody, err = c.GetRawData()
	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	if len(bytesBody) <= 0 {
		log.Println(`empty body`)
		err = errors.New(`empty body`)
		return
	}

	buffer := bytes.NewBuffer(bytesBody)
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		gz, tmpErr := gzip.NewReader(buffer)
		if tmpErr != nil {
			err = tmpErr
			log.Printf("gzip fail %v \n", err)
			return
		}
		buf := new(bytes.Buffer)
		n, tmpErr := buf.ReadFrom(gz)
		if tmpErr != nil {
			err = tmpErr
			log.Printf("gzip reader fail %v \n", err)
			return
		}
		if n <= 0 {
			err = errors.New("empty body")
			return
		}
		bytesBody = buf.Bytes()
		return
	case "deflate":
		def := flate.NewReader(buffer)
		defer def.Close()
		buf := new(bytes.Buffer)
		n, tmpErr := buf.ReadFrom(def)
		if tmpErr != nil {
			err = tmpErr
			log.Printf("deflate reader fail %v \n", err)
			return
		}
		if n <= 0 {
			err = errors.New("empty body")
			return
		}
		bytesBody = buf.Bytes()
		return
	}
	return
}
