package main

import "fmt"

func main() {
	fmt.Printf("Initializing certificates...\n")
	serverTLSConf, clientTLSConf, caPEM, err := certsetup()
	if err := nil {
		panic(err)
	}
}

s := Server{
	ServerTLSConf: serverTLSConf,
	ClientTLSConf: clientTLSConf,
	CaPEM: caPEM,
}

go func(){
	handler := http.NewServeMux()

	handler.HandleFunc("/ca.pem", s.getCA)
	fmt.Printf("Starintg localhost http server on :8080 with ca.pem endpoint\n")
	err = http.ListenAndServe("localhost:8080", handler)
}