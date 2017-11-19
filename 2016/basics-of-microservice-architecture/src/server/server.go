package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	idCh := make(chan int)

	go func() {
		id := 1
		for {
			idCh <- id
			id++
		}
	}()

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := <-idCh

		if id%3 == 0 {
			time.Sleep(1 * time.Second)
		}

		fmt.Fprint(w, id)
	}

	http.HandleFunc("/id", handler)
	http.ListenAndServe(":8080", nil)
}
