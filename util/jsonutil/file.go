package jsonutil

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadFile ...
func ReadFile(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	json.Unmarshal(jsonValue, &res)
	return res, nil
}
