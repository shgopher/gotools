package gotools

import "testing"

func TestReadFile(t *testing.T) {
	t.Log(ReadFile("your path"))
}

func TestRead(t *testing.T) {
	t.Log(Read("your path"))
}