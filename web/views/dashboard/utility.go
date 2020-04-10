package dashboard

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("../web.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "Web INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Fatalln(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "0.0.1"
}
