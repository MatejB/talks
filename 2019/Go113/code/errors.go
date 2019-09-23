package main

import (
	"errors"
	"fmt"
	"time"
)

var SentinelErr = errors.New("Sentinel Error")

type myError struct {
	Err  error
	Time time.Time
}

func (err myError) Error() string {
	return fmt.Sprintf("%s @ %s", err.Err, err.Time)
}
func (err myError) Unwrap() error {
	return err.Err
}

func foo() error {
	return myError{Err: SentinelErr, Time: time.Now()}
}

func bar() error {
	err := foo()
	if err != nil {
		return fmt.Errorf("foo failed: %w", err)
	}
	return nil
}

func main() {
	if err := bar(); err != nil {
		fmt.Println(err)

		if errors.Is(err, SentinelErr) {
			fmt.Println("Is a Sentinel Error")
		}

		var myErr myError
		if errors.As(err, &myErr) {
			fmt.Println("Error time is", myErr.Time.Format(time.RFC3339))
		}
	}
}
