package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	code := http.StatusOK
	w.WriteHeader(http.StatusOK)
	msg := make(map[string]string)
	msg["code"] = "200"
	msg["msg"] = "this is Healthz!"
	res, err := json.Marshal(msg)
	if err != nil {
		log.Println("ERROR:  %V", err)
		fmt.Fprint(w, "系统错误请重新")
	}
	fmt.Fprint(w, string(res))
	setLog(r, code, string(res))
}

func Other(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is "+r.URL.String())
}

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Other)
	http.ListenAndServe(":80", nil)
}

//设置head头
func setHeader(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "none"
	}
	w.Header().Add("VERSION", version)
}

//记录日志  客户端 IP，HTTP 返回码
func setLog(r *http.Request, code int, res string) {
	log.Println(fmt.Sprintf("%v   %v  ip: %v  Response: %v   %v", r.Method, r.URL.String(), r.RemoteAddr, code, res))
}
