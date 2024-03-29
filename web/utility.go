package web

import (
	"encoding/json"
	"fmt"
	m "kRtrima/plugins/database/mongoDB/models"
	"log"
	"net/http"
	"os"
	"strings"
)

// Configuration for the web server
type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var lm = &m.Msg

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	file, err := os.OpenFile("kRtrima_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "Web: ", log.Ldate|log.Ltime|log.Lshortfile)
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("web/config/config.json")
	defer file.Close()
	if err != nil {
		logger.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatalln("Cannot get configuration from file", err)
	}
}

// Convenience function to redirect to the error message page
func errormessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

//to override the post method to follow the restful convention
func methodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only act on POST requests.
		if r.Method == "POST" {

			// Look in the request body and headers for a spoofed method.
			// Prefer the value in the request body if they conflict.
			method := r.PostFormValue("_method")
			if method == "" {
				method = r.Header.Get("X-HTTP-Method-Override")
			}

			// Check that the spoofed method is a valid HTTP method and
			// update the request object accordingly.
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
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
