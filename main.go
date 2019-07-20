package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serverHome(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL)
	http.ServeFile(w, r, "home.html")
}

func main() {
	fmt.Println("run main")
	flag.Parse()
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/", serverHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	checkErr(err)
}
