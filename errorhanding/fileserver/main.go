package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 统一错误处理
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if _, ok := r.(error); ok {
				log.Printf("Painc:%s", r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		var err = f(w, r)
		var status = http.StatusOK
		if err != nil {
			log.Printf("Error handling request:%s", err.Error())

			if e, ok := err.(UserError); ok {
				http.Error(w, e.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				status = http.StatusNotFound
			case os.IsPermission(err):
				status = http.StatusForbidden
			default:
				status = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func main() {
	http.HandleFunc("/", errorHandler(handleFileList()))
	log.Println("start server at port:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

const prefix = "/list/"

// 路由对应的处理函数
func handleFileList() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		if strings.Index(r.URL.Path, prefix) == -1 {
			return CanSeeError("path must start with " + prefix)
		}
		var path = r.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(w, string(bytes))
		return nil
	}
}

type UserError interface {
	error
	Message() string
}

type CanSeeError string

func (e CanSeeError) Error() string {
	return e.Message()
}

func (e CanSeeError) Message() string {
	return string(e)
}
