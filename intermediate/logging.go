package intermediate

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message.")

	log.SetPrefix("INFO:")
	log.Println("This is an info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with date and time.")

	infoLogger.Println("This os an info message")
	warnLogger.Println("This is an warning message.")
	errorLogger.Println("This is an error message.")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	defer file.Close()

	infoLogger1 := log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message.")
	warnLogger1.Println("This is a warning message.")
	infoLogger1.Println("This is an info message.")
	errorLogger1.Println("This is an error.")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
)
