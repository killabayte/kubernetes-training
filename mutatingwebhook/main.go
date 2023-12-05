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