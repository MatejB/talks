package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func getID() {
	defer wg.Done()

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

var wg sync.WaitGroup

func main() {
	// START OMIT
	reqSlots := make(chan struct{}, 3)
	var token struct{}

	go func() {
		for {
			<-reqSlots
			go getID()
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case reqSlots <- token:
			wg.Add(1)
		default:
			fmt.Printf("All slots full, dropping req %d\n", i)
			time.Sleep(1 * time.Millisecond)
		}
	}
	// END OMIT

	wg.Wait()
}
