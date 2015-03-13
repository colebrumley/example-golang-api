package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen   = flag.String("listen", ":80", "Address to listen on")
	tls      = flag.Bool("tls", false, "Use TLS (requires -cert and -key)")
	cert     = flag.String("cert", "", "TLS cert file")
	key      = flag.String("key", "", "TLS key file")
	password = flag.String("password", "", "Enable basic authentication")
)

func main() {
	flag.Parse()
	mux := http.NewServeMux()

	if len(*password) > 0 {
		log.Println("[notice] Using HTTP basic authentication")
		mux.HandleFunc("/", BasicAuth(*password, DefaultHandler))
	} else {
		mux.HandleFunc("/", DefaultHandler)
	}
	if *tls {
		if *listen == ":80" {
			*listen = ":443"
		}
		log.Println("[notice] Listening on", *listen, "over TLS")
		log.Fatal(http.ListenAndServeTLS(*listen, *cert, *key, mux))
	} else {
		log.Println("[notice] Listening on", *listen)
		log.Fatal(http.ListenAndServe(*listen, mux))
	}
}

func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("It works!"))
}
