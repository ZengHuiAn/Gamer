package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/static",serveTemplate)
	log.Println("Listening...")
	http.ListenAndServe(":13000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// content:= len(r.Form)
	print("querys",r.Form)
	print("proto:",r.Proto)
	//w.WriteHeader(http.StatusGatewayTimeout)
	// w.Write([]byte(string(content)))
}
