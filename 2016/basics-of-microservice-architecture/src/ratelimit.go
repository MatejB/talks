package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getID() {
	client := http.Client{}

	res, err := client.Get("http://localhost:8080/id")
	if err != nil {
		log.Fatalln(err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res.Body.Close()

	id, err := strconv.Atoi(buf.String())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(id)
}

func main() {
	// START OMIT
	reqLimit := time.Second / 2

	throttle := time.Tick(reqLimit)

	for i := 0; i < 10; i++ {
		<-throttle
		getID()
	}
	// END OMIT
}
