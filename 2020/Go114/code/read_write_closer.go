package main

import "io"

type ReadWriteCloser interface {
	io.ReadCloser
	io.WriteCloser
}

func main() {}
