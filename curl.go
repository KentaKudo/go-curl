package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Curl struct {
	Method  string
	URL     string
	Verbose bool
}

func (c *Curl) request(params map[string]string) error {
	resp, err := http.Get(c.URL)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
