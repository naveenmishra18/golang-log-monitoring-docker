package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {

	file, err := os.OpenFile("/var/log/app/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	logger.Out = file
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Info("INFO: Service accessed")
	logger.Warn("WARN: Warning detected")
	logger.Debug("DEBUG: Debugging information")
	logger.Error("ERROR: An error occurred")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logs generated successfully!")
}

func main() {
	http.HandleFunc("/", handler)
	logger.Info("Starting the server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server failed to start:", err)
	}
}
