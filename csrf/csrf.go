package csrf

import (
	"errors"
	"log"

	"github.com/jagochiu/utils/jrsa"
)

/*
CSRF -
*/
type CSRF struct {
	ID  string `json:"_csrf,omitempty" bson:"_id,omitempty"`
	Pub string `json:"pub,omitempty" bson:"pub,omitempty"`
	Pri string `json:"pri,omitempty" bson:"pri,omitempty"`
}

/*
RsaDecryption -
*/
func (c *CSRF) Decrypt(data string) (resp string, err error) {
	resp = ``
	if len(data) <= 0 {
		err = errors.New(`empty data`)
		return
	}
	if len(c.Pri) <= 0 {
		err = errors.New(`empty private key`)
		return
	}
	r := jrsa.JRSA{}
	err = r.Set(c.Pri, jrsa.Private)
	if err != nil {
		log.Printf("%v \n", err)
	}
	resp, err = r.PriDecrypt(data)
	if err != nil {
		log.Printf("[PRIVATE DECRYPTED] %v \n", err)
	}
	return
}

/*
Encrypt -
*/
func (c *CSRF) Encrypt(data string) (resp string, err error) {
	resp = ``
	if len(data) <= 0 {
		err = errors.New(`empty data`)
		return
	}
	if len(c.Pub) <= 0 {
		err = errors.New(`empty public key`)
		return
	}
	r := jrsa.JRSA{}
	err = r.Set(c.Pub, jrsa.Public)
	if err != nil {
		log.Printf("%v \n", err)
	}
	resp, err = r.PubEncrypt(data)
	if err != nil {
		log.Printf("[PUBLIC ENCRYPTED] %v \n", err)
	}
	return
}
