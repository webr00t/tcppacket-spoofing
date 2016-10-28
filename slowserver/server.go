package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...=" + os.Getenv("PORT"))
	err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	sleep, err := time.ParseDuration(req.URL.Query().Get("sleep"))
	if err != nil {
		sleep, _ = time.ParseDuration("5s")
	}
	fmt.Fprintln(res, "Slept for ", sleep)

	time.Sleep(sleep)
}