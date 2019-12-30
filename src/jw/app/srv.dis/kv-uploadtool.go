package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main2() {
	consulUrl := os.Args[1]
	fileName := os.Args[2]
	//keyUrl := os.Args[3]

	f, err := os.Open(fileName)
	//content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("consulUrl: %v", consulUrl)
	req, err := http.NewRequest("PUT", consulUrl, f)
	if err != nil {
		log.Fatal("NewRequest error:" + err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Do request error:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Do request error:" + err.Error())
	}

	log.Printf("consul response: %v", string(body))
}