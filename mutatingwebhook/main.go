package main

import "fmt"

func main() {
	fmt.Printf("Initializing certificates...\n")
	serverTLSConf, clientTLSConf, caPEM, err := certsetup()

}
