package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello k8s test go 触发器 222"))
}
