package jrsa

import (
	"testing"
)

func TestCreateKeys(t *testing.T) {
	r := JRSA{}
	err := r.CreateKeys(2048)
	if err != nil {
		t.Error(err.Error())
	}

	r.Print()
	// publicBuf := bytes.NewBufferString("")
	// privateBuf := bytes.NewBufferString("")
	// if err := CreateKeysO(publicBuf, privateBuf, 2048); err == nil {
	// 	if _, err := NewJRSA(publicBuf.Bytes(), privateBuf.Bytes()); err == nil {
	// 		println(publicBuf.String())
	// 		println(`-----------------------------------`)
	// 		println(privateBuf.String())
	// 	} else {
	// 		log.Printf("GenerateRSA : %v \n", err)
	// 	}
	// } else {
	// 	log.Printf("GenerateRSA Key : %v \n", err)
	// }
}
