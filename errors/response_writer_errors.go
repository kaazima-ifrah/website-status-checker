package errors

import "log"

func HandleResponseWriterError(funcName string, err error) {
	log.Fatal("Error while adding the", funcName, "response to response writer! ", err)
	return
}
