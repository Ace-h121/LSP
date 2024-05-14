package main

import (
	"bufio"
	"log"
	"os"

	"Github.com/LSP/rpc"
)



func main(){
	logger := getLogger("/home/ace/Documents/Programming/Networking/LSP/log.txt")
	logger.Println("Logger Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan(){
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any){
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger{
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil{
		panic("better file pls")
	}
	return log.New(logfile, "[LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
