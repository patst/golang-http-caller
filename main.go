package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(argsWithoutProg[0])
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
