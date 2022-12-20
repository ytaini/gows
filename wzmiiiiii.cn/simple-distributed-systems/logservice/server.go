package logservice

import (
	"io"
	"log"
	"net/http"
	"os"
)

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

var logger *log.Logger

func Init(destPath string) {
	logger = log.New(fileLog(destPath), "", 0)
}

func RegisterHandlerFunc() {
	http.HandleFunc("/log", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		msgBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		logger.Print(string(msgBytes))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
