package utils

import (
	"net/http"
)

//SendResponse sends the data in form of JSON
func SendResponse(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("Contant-type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}

//SendError sends an error
func SendError(writer http.ResponseWriter, status int, errorMsg string) {
	data := []byte(errorMsg)
	writer.Header().Set("Contant-type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
