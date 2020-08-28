package jsonutil

import (
	"encoding/json"
	"io/ioutil"
)

// ReadFile ...
var ReadFile = func(path string, in interface{}) error {
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
