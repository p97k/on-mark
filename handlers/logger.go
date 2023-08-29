package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Logger struct {
	log *log.Logger
}

func (logger *Logger) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("yup, you are good to go!")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Something Went Wrong!", http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(response, "Data: %s\n", data)
	if err != nil {
		http.Error(response, "Something Went Wrong!", http.StatusBadRequest)
		return
	}
}

func NewLogger(log *log.Logger) *Logger {
	return &Logger{log}
}
