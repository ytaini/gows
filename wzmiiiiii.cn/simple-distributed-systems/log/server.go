package log

import (
	"io"
	"net/http"
)

// 日志服务的目的:
// 首先它是一个web服务,它可以接收post请求,然后将post请求中的内容,写入log中.

// 将日志写入文件系统
// fileLog 就是写入文件的路径
//type fileLog string
//
//func (fl fileLog) Write(data []byte) (int, error) {
//	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
//	if err != nil {
//		return 0, err
//	}
//	defer f.Close()
//	return f.Write(data)
//}

func Run(destination string) {
	log = InitLogger(destination)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(msg string) {
	log.Infof("%v\n", msg)
}
