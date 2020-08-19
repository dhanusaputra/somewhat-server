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
func ReadFile(path string) (interface{}, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var res interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
