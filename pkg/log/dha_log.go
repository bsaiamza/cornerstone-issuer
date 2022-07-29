package log

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// var (
// 	DHAInfo    *log.Logger
// 	DHAError   *log.Logger
// 	DHAWarning *log.Logger

// 	DHAServerInfo    *log.Logger
// 	DHAServerError   *log.Logger
// 	DHAServerWarning *log.Logger
// )

// func init() {
// 	dhaLogFile, err := os.OpenFile("dha_query.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		log.Printf("Failed to create dha_query log file: %s", err.Error())
// 	}
// 	serverLogFile, err := os.OpenFile("dha_query_server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		log.Printf("Failed to create dha_query_server log file: %s", err.Error())
// 	}

// 	dmw := io.MultiWriter(os.Stdout, dhaLogFile)
// 	smw := io.MultiWriter(os.Stdout, serverLogFile)

// 	Info = log.New(dmw, "[DHA][INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
// 	Error = log.New(dmw, "[DHA][ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
// 	Warning = log.New(dmw, "[DHA][WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)

// 	ServerInfo = log.New(smw, "[DHA][INFO]: \t", log.Ldate|log.Ltime|log.Lshortfile)
// 	ServerError = log.New(smw, "[DHA][ERROR]: \t", log.Ldate|log.Ltime|log.Lshortfile)
// 	ServerWarning = log.New(smw, "[DHA][WARN]: \t", log.Ldate|log.Ltime|log.Lshortfile)
// }
