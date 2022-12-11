package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	err := SetupLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

func SetupLogger() error {
	file, err := os.OpenFile("./logtest.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	} else {
		defer file.Close()
	}
	log.SetOutput(file)
	return nil
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
