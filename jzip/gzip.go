package jzip

import (
	"bytes"
	"compress/gzip"
	"log"
)

/*
GZipString -
*/
func GZipString(data string, level int) string {
	return string(GZip([]byte(data), level))
}

/*
GZip -
*/
func GZip(data []byte, level int) []byte {
	var b bytes.Buffer
	if gz, err := gzip.NewWriterLevel(&b, level); err == nil {
		if _, err := gz.Write(data); err != nil {
			log.Fatal(err)
		}
		if err := gz.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return b.Bytes()
}
