package main

import "fmt"

// Writer interface
type Writer interface {
	Write(string)
}

// ConsoleWriter implementation
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(msg string) {
	fmt.Println("Writing to Console : ", msg)
}

// XMLWriter implementation
type XMLWriter struct{}

func (xw XMLWriter) Write(msg string) {
	fmt.Println("Writing to XML : ", msg)
}

// Logger implementation
type Logger struct {
	Message string
}

// Log implementation
func (logger Logger) Log(writer Writer) {
	writer.Write(logger.Message)
}
