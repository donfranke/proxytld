package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)
// generic error handler
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
    tldfile := flag.String("t", "", "Location/Name of TLD File")
    urlfile := flag.String("p", "", "Location/Name of Proxy File")
    flag.Parse()

    // make sure all flags are provided by user
    if *tldfile == "" || *urlfile == "" {
        log.Fatal("EXECUTION HALTED: Not enough arguments supplied\n\n" + showUsage())
    }
    log.Print("Started")

	tlds := loadTLDs(*tldfile)
	urls := loadURLs(*urlfile)

	// iterate urls and compare against TLDs
	for _, url := range urls {
		url := strings.Replace(url, "\"", "", -1)

		s := strings.Split(url, ".")
		thistld := s[len(s)-1]
		thistld = strings.Replace(thistld, "\"", "", -1)
		thistld = "." + thistld

		for _, tld := range tlds {
			tld = strings.Replace(tld, "\"", "", -1)

			if thistld == tld {
				fmt.Printf("MATCH FOUND: %s\n",url)
			}
			//fmt.Printf("This TLD: %s, URL: %s, TLD: %s\n",thistld2, url2,tld2)
		}
	}
	log.Print("Finished")

}

// load offending tlds into memory
func loadTLDs(tf string) []string {
	// load tlds file into memory
	var t []string
	dat, err := ioutil.ReadFile(tf)
	check(err)
	t = strings.Split(string(dat), "\n")
	fmt.Printf("  TLDs loaded: %d\n", len(t))
	
	return t
}

// load urls (such as from proxy logs) into memory
func loadURLs(uf string) []string {
	var u []string	
	dat, err := ioutil.ReadFile(uf)
	check(err)
	u = strings.Split(string(dat), "\n")
	fmt.Printf("  URLs loaded: %d\n", len(u))

	return u
}

func showUsage() string {
	var message string
	message = "\t-t = path/file of TLD file\n"
	message += "\t-p = path/file of Proxy file\n"

	return message
}