package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(req *http.Request, data interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(data)
	PanicHandler(err)
}

func WriteEncodeJson(wr http.ResponseWriter, response interface{}) {
	wr.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(wr)
	err := encoder.Encode(response)
	PanicHandler(err)
}
