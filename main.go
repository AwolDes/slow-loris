package main

import (
	"fmt"
	"flag"
	"os"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: slow-loris -u https://test.com -p 443 -c 500")
	os.Exit(1)
}

func main() {
	// invoked like: slow-loris {url/ip} {port} {connections}

	var defaultPort int = 80
	var defaultConnections int = 100

	url := flag.String("u", "nil", "This is the url or IP to attack")
	port := flag.Int("p", defaultPort, "This is the port that is attacked (default is 80)")
	connections := flag.Int("c", defaultConnections, "This is the number of concurrent connections (default is 100)")

	flag.Parse()

	
	if flag.NArg() == 0 { 
		printUsage()
	}

	if *url == "nil" {
		panic("URL is required")
	}

	fmt.Println(*url, *port, *connections)
}