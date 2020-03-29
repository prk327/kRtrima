package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"kRtrima/plugins/database/mongoDB"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	file, err := os.OpenFile("web/log/kRtrima.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
    loadConfig()
	initializeDB()
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

func initializeDB() {
	*lm, mongoDB.DB = mongoDB.Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	fmt.Println(mongoDB.Msg)

	//    conect to collection
	*lm, mongoDB.Collection = mongoDB.Cnt_Collection("Thread", mongoDB.DB)
	fmt.Println(mongoDB.Msg)
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
//func session(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) (sess data.Session, err error) {
//	cookie, err := request.Cookie("_cookie")
//	if err == nil {
//		sess = data.Session{Uuid: cookie.Value}
//		if ok, _ := sess.Check(); !ok {
//			err = errors.New("Invalid session")
//		}
//	}
//	return
//}

// parse HTML templates
// pass in a list of file names, and get a template
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var t *template.Template
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/ui/template/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(writer, "layout", data)
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
