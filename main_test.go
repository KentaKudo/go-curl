package main

import (
	"bytes"
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
