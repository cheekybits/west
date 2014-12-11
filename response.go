package west

import (
	"io/ioutil"
	"net/http"
)

// Resp represents a response.
type Resp struct {
	*http.Response
	parsedBody bool
	body       []byte
}

// BodyBytes gets the bytes from the response body.
func (r *Resp) BodyBytes() []byte {
	if !r.parsedBody {
		r.parsedBody = true
		var err error
		if r.body, err = ioutil.ReadAll(r.Body); err != nil {
			return []byte(err.Error())
		}
	}
	return r.body
}

// BodyString gets the body as a string.
func (r *Resp) BodyString() string {
	return string(r.BodyBytes())
}

// BodyMap gets the body as a map[string]interface{}.
func (r *Resp) BodyMap() map[string]interface{} {
	var obj map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &obj); err != nil {
		panic("BodyMap failed: " + err.Error())
	}
	return obj
}

// BodyMapSlice gets the body as a []map[string]interface{}.
func (r *Resp) BodyMapSlice() []map[string]interface{} {
	var objs []map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &objs); err != nil {
		panic("BodyMapSlice failed: " + err.Error())
	}
	return objs
}
