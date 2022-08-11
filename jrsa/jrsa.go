package jrsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"strings"
)

type KeyType string

const (
	Private KeyType = "PRIVATE"
	Public          = "PUBLIC"
)

/*
JRSA -
*/
type JRSA struct {
	PubKey *rsa.PublicKey
	PriKey *rsa.PrivateKey
}

/*
PriDecrypt -
*/
func (r *JRSA) Set(k string, t KeyType) error {
	keyBlock, err := stringCoverKeyBlock(k, t)
	if err != nil {
		log.Printf("%v \n", err)
		return err
	}
	switch t {
	case Private:
		priKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
		if err != nil {
			log.Printf("%v \n", err)
			return err
		}
		r.PriKey = priKey.(*rsa.PrivateKey)
	case Public:
		pubKey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
		if err != nil {
			log.Printf("%v \n", err)
			return err
		}
		r.PubKey = pubKey.(*rsa.PublicKey)
	}
	return nil
}

/*
------------------------------------------------------------------------------------------------------
*/

/*
CreateKeys -
*/
func (r *JRSA) CreateKeys(keyLength int) (err error) {
	priKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	r.PriKey = priKey
	r.PubKey = &priKey.PublicKey
	return
}

/*
Print -
*/
func (r *JRSA) Print() {
	keyStr, err := r.String(Private)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	println(keyStr)
	keyStr, err = r.String(Public)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	println(keyStr)
}

/*
ExportRSA -
*/
func (r *JRSA) String(keyType KeyType) (rs string, err error) {
	var bytes []byte
	rs = ``
	switch keyType {
	case Private:
		bytes, err = x509.MarshalPKCS8PrivateKey(r.PriKey)
	case Public:
		bytes, err = x509.MarshalPKIXPublicKey(r.PubKey)
	default:
		return ``, errors.New(``)
	}
	if err != nil {
		return
	}
	pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  fmt.Sprintf(`RSA %s KEY`, keyType),
			Bytes: bytes,
		},
	)
	return string(pem), nil
}

/*
------------------------------------------------------------------------------------------------------
*/

/*
PubEncrypt -
*/
func (r *JRSA) PubEncrypt(data string) (string, error) {
	partLen := r.PubKey.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bts, err := rsa.EncryptPKCS1v15(rand.Reader, r.PubKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bts)
	}
	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

/*
PriDecrypt -
*/
func (r *JRSA) PriDecrypt(encrypted string) (string, error) {
	partLen := r.PriKey.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(encrypted)
	chunks := split([]byte(raw), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.PriKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	return buffer.String(), err
}

/*
------------------------------------------------------------------------------------------------------
*/

/*
PubEncrypt -
*/
func PubEncrypt(data, key string) (rs string, err error) {
	buffer := bytes.NewBufferString("")
	pubBlock, err := stringCoverKeyBlock(key, Public)
	if err != nil {
		log.Printf("%v \n", err)
	}
	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		log.Printf("%v \n", err)
	}
	r := JRSA{PubKey: pubKey.(*rsa.PublicKey)}
	partLen := r.PubKey.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)
	for _, chunk := range chunks {
		bts, err := rsa.EncryptPKCS1v15(rand.Reader, r.PubKey, chunk)
		if err != nil {
			return ``, err
		}
		buffer.Write(bts)
	}
	rs = base64.RawURLEncoding.EncodeToString(buffer.Bytes())
	return
}

/*
PriDecrypt -
*/
func PriDecrypt(data, key string) (rs string, err error) {
	priBlock, err := stringCoverKeyBlock(key, Private)
	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	priKey, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
	if err != nil {
		log.Printf("%v \n", err)
	}
	r := JRSA{PriKey: priKey.(*rsa.PrivateKey)}
	buffer := bytes.NewBufferString("")
	raw, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	partLen := r.PriKey.N.BitLen() / 8
	chunks := split([]byte(raw), partLen)
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.PriKey, chunk)
		if err != nil {
			return ``, err
		}
		buffer.Write(decrypted)
	}
	rs = buffer.String()
	return
}

/*
------------------------------------------------------------------------------------------------------
*/

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

func stringCoverKeyBlock(keyString string, keyType KeyType) (*pem.Block, error) {
	if len(keyString) <= 64 {
		return nil, errors.New("key string length error")
	}
	temp := string("-----BEGIN " + keyType + " KEY-----\n")
	for i, char := range keyString {
		temp += fmt.Sprintf(`%c`, char)
		if (i+1)%64 == 0 {
			temp += "\n"
		}
	}
	temp += string("\n-----END " + keyType + " KEY-----")
	buff := bytes.NewBufferString(temp)
	block, _ := pem.Decode(buff.Bytes())
	if block == nil {
		return nil, errors.New(`key empty`)
	}
	return block, nil
}

func SafeBase64EncodeURL(data string) string {
	return strings.ReplaceAll(strings.ReplaceAll(data, `+`, `-`), `/`, `_`)
}
func SafeBase64DecodeURL(data string) string {
	return strings.ReplaceAll(strings.ReplaceAll(data, `-`, `+`), `_`, `/`)
}
