package west

import (
	"encoding/json"
)

type MarshalFunc func(v interface{}) ([]byte, string, error)

// Marshal marshals the object into bytes.
var Marshal MarshalFunc = func(v interface{}) ([]byte, string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, "", err
	}
	return b, "application/json; charset=utf-8", nil
}

// Unmarshal unmsrahsls from the bytes into the object.
var Unmarshal = json.Unmarshal
