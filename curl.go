package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Curl struct {
	Client  *http.Client
	URL     string
	Verbose bool
	Headers Headers
}

func NewCurl(url string, verbose bool, headers Headers) *Curl {
	return &Curl{
		Client:  &http.Client{},
		URL:     format(url),
		Verbose: verbose,
		Headers: headers,
	}
}

func (c *Curl) Get(params map[string]string) (string, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return "", err
	}

	for _, header := range c.Headers {
		h := strings.Split(header, ": ")
		if len(h) != 2 {
			continue
		}

		req.Header.Add(h[0], h[1])
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Curl) Post(data string) (string, error) {
	req, err := http.NewRequest("POST", c.URL, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	for _, header := range c.Headers {
		h := strings.Split(header, ": ")
		if len(h) != 2 {
			continue
		}

		req.Header.Add(h[0], h[1])
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func format(url string) string {
	if !strings.HasSuffix(url, "http") {
		return "http://" + url
	}
	return url
}
