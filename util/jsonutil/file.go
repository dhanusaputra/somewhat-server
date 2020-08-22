package jsonutil

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type ioReader interface {
	io.Reader
}

// ReadFile ...
func ReadFile(path string, in interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &in)
	if err != nil {
		return err
	}
	return nil
}
