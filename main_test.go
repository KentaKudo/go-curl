package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestStatusCode(t *testing.T) {
	sut := &CLI{outStream: new(bytes.Buffer), errStream: new(bytes.Buffer)}
	args := strings.Split(`-v -X "GET" https://google.com`, " ")
	if status := sut.Run(args); status != ExitCodeOK {
		t.Errorf("got %d, want %d", status, ExitCodeOK)
	}
}

func TestStatusCodeError(t *testing.T) {
	sut := &CLI{outStream: new(bytes.Buffer), errStream: new(bytes.Buffer)}
	args := strings.Split(`-unknown-option https://google.com`, " ")
	if status := sut.Run(args); status != ExitCodeError {
		t.Errorf("got %d, want %d", status, ExitCodeError)
	}
}

func TestRun(t *testing.T) {
	url := "example.com"
	outStream := new(bytes.Buffer)
	sut := &CLI{outStream: outStream, errStream: new(bytes.Buffer)}
	args := strings.Split(fmt.Sprintf("%s", url), " ")

	sut.Run(args)
	curl := exec.Command("curl", url)
	out, err := curl.Output()
	if err != nil {
		t.Errorf(`"curl %s" failed! Please check the network connection: %s`, url, err)
	}
	want := string(out)
	if outStream.String() != want {
		t.Errorf("got %q, want %q", outStream.String(), want)
	}
}
