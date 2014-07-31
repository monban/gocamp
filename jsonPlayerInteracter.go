package gocamp

import "encoding/json"

type JsonPlayerInteracter struct {
}

func (self JsonPlayerInteracter) ProcessMessage(message string) (result string) {
	var data interface{}
	_ = json.Unmarshal([]byte(message), &data)
	return ""
}
