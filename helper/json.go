package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfErr(err)
}

func WriteToResponseBody(write http.ResponseWriter, response any) {
	write.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	PanicIfErr(err)
}
