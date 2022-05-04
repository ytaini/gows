package gobasic

/*
   大多数情况下,当你的http响应失败时,resp变量将为nil,而err 变量将是non-nil.
   然而,当你得到一个重定向的错误时,两个变量都是non-nil.这意味着你最后依然会内存泄露.
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Eg81() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
