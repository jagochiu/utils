package webtools

import (
	"encoding/json"
	"log"
)

/*
Respond -
*/
type Respond struct {
	Status    int         `json:"status" bson:"status"`
	Data      interface{} `json:"data" bson:"data"`
	TimeStamp int64       `json:"timestamp" bson:"timestamp"`
	Created   int64       `json:"created,omitempty" bson:"created,omitempty"`
	Message   string      `json:"message,omitempty" bson:"message,omitempty"`
	Error     string      `json:"error,omitempty" bson:"error,omitempty"`
	Path      string      `json:"path,omitempty" bson:"path,omitempty"`
}

/*
RespondParser -
*/
func RespondParser(data interface{}) (r Respond, err error) {
	switch data.(type) {
	case string:
		err = json.Unmarshal([]byte(data.(string)), &r)
		if err != nil {
			log.Fatalf("[JSON Unmarshal]: %d \n", err)
		}
	case []byte:
		err = json.Unmarshal(data.([]byte), &r)
		if err != nil {
			log.Fatalf("[JSON Unmarshal]: %d \n", err)
		}
	default:
		tmp, err := json.Marshal(data)
		if err != nil {
			log.Printf("[JSON Marshal]: %d \n", err)
		}
		err = json.Unmarshal(tmp, &r)
		if err != nil {
			log.Printf("[JSON Unmarshal]: %d \n", err)
		}
	}
	return
}

/*
Parser -
*/
func (r *Respond) Parser(data string) error { return json.Unmarshal([]byte(data), r) }
