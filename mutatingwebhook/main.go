package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Initializing certificates...\n")
	serverTLSConf, clientTLSConf, caPEM, err := certsetup()

	if err != nil {
		panic(err)
	}

	s := Server{
		ServerTLSConf: serverTLSConf,
		ClientTLSConf: clientTLSConf,
		CaPEM:         caPEM,
	}

	go func() {
		handler := http.NewServeMux()

		handler.HandleFunc("/ca.pem", s.getCA)
		fmt.Printf("Starintg localhost http server on :8080 with ca.pem endpoint\n")
		err = http.ListenAndServe("localhost:8080", handler)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Starting TLS server on :8443\n")
	handler := http.NewServeMux()
	handler.HandleFunc("/webhook", s.postWebhook)

	https := &http.Server{
		Addr:      ":8443",
		TLSConfig: serverTLSConf,
		Handler:   handler,
	}

	log.Fatal(https.ListenAndServeTLS("", ""))

}
