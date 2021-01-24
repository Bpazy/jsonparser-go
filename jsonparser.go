package jsonparser

import (
	"encoding/json"
)

func Parse(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
