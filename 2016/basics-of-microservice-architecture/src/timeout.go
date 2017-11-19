package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// START OMIT
		client := http.Client{
			Timeout: time.Duration(50 * time.Millisecond),
		}

		res, err := client.Get("http://localhost:8080/id")
		if err != nil {
			fmt.Println("TIMEOUT")
			continue
		}

		io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		// END OMIT

		fmt.Println("")
	}
}
