package jfile

import (
	"io/ioutil"
	"log"
	"os"
)

func Read(path string) ([]byte, error) {
	data, err := os.Open(path)
	if err != nil {
		log.Printf("open sql file - %v \n", err)
		return nil, err
	}
	return ioutil.ReadAll(data)
}
