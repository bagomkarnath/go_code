package main

func main() {
	var logger Logger = Logger{Message: "Test log message"}
	var cw ConsoleWriter
	logger.Log(cw)
	var xw XMLWriter
	logger.Log(xw)
}
