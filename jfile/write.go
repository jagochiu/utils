package jfile

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func Write(data string, path string) error { return ioutil.WriteFile(path, []byte(data), 0777) }
func WriteObj(data interface{}, path string) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("JSON Parsing Object - %v \n", err)
		return err
	}
	err = ioutil.WriteFile(path, jsonBytes, 0777)
	if err != nil {
		log.Printf("Write games JSON file - %v \n", err)
	}
	return nil
}
