package log

import (
	"io"
	"log"
	"os"
)

var (
	// Info ...
	Info *log.Logger
	// Error ...
	Error *log.Logger
	// Warning ...
	Warning *log.Logger

	// ServerInfo ...
	ServerInfo *log.Logger
	// ServerError ...
	ServerError *log.Logger
	// ServerWarning ...
	ServerWarning *log.Logger
)

func init() {
	issuerLogFile, err := os.OpenFile("cornerstone_issuer.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Print("Error: Failed to create cornerstone_issuer log file:", err)
	}
	serverLogFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Print("Error: Failed to create server log file:", err)
	}

	imw := io.MultiWriter(os.Stdout, issuerLogFile)
	smw := io.MultiWriter(os.Stdout, serverLogFile)

	Info = log.New(imw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(imw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(imw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)

	ServerInfo = log.New(smw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerError = log.New(smw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerWarning = log.New(smw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)
}
