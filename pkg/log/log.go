package log

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger

	ServerInfo    *log.Logger
	ServerError   *log.Logger
	ServerWarning *log.Logger

	DHAInfo    *log.Logger
	DHAError   *log.Logger
	DHAWarning *log.Logger
)

func init() {
	issuerLogFile, err := os.OpenFile("cornerstone_issuer.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Failed to create cornerstone_issuer log file: %s", err.Error())
	}
	serverLogFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Failed to create server log file: %s", err.Error())
	}
	dhaLogFile, err := os.OpenFile("dha_query.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Failed to create dha_query log file: %s", err.Error())
	}

	imw := io.MultiWriter(os.Stdout, issuerLogFile)
	smw := io.MultiWriter(os.Stdout, serverLogFile)
	dmw := io.MultiWriter(os.Stdout, dhaLogFile)

	Info = log.New(imw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(imw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(imw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)

	ServerInfo = log.New(smw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerError = log.New(smw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	ServerWarning = log.New(smw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)

	DHAInfo = log.New(dmw, "[INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	DHAError = log.New(dmw, "[ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
	DHAWarning = log.New(dmw, "[WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)
}
