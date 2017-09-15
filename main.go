package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func CheckUrl(url string) [20]byte {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return sha1.Sum(b)
}

func main() {
	fmt.Printf("%x\n\n", CheckUrl(os.Args[1]))
}
