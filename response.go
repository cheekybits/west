package west

import (
	"io/ioutil"
	"net/http"
)

// Response represents a Responseonse.
type Response struct {
	*http.Response
	parsedBody bool
	body       []byte
}

// BodyBytes gets the bytes from the Responseonse body.
func (r *Response) BodyBytes() []byte {
	if !r.parsedBody {
		defer r.Body.Close()
		r.parsedBody = true
		var err error
		if r.body, err = ioutil.ReadAll(r.Body); err != nil {
			return []byte(err.Error())
		}
	}
	return r.body
}

// BodyString gets the body as a string.
func (r *Response) BodyString() string {
	return string(r.BodyBytes())
}

// BodyObj gets the body as an interface{}.
func (r *Response) BodyObj() interface{} {
	var obj interface{}
	if err := Unmarshal(r.BodyBytes(), &obj); err != nil {
		panic("BodyObj failed: " + err.Error())
	}
	return obj
}

// BodyMap gets the body as a map[string]interface{}.
func (r *Response) BodyMap() map[string]interface{} {
	var obj map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &obj); err != nil {
		panic("BodyMap failed: " + err.Error())
	}
	return obj
}

// UnmarshalBody unmarshals the body into the given object.
func (r *Response) UnmarshalBody(obj interface{}) {
	if err := Unmarshal(r.BodyBytes(), &obj); err != nil {
		panic("UnmarshalBody failed: " + err.Error())
	}
}

// BodyMapSlice gets the body as a []map[string]interface{}.
func (r *Response) BodyMapSlice() []map[string]interface{} {
	var objs []map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &objs); err != nil {
		panic("BodyMapSlice failed: " + err.Error())
	}
	return objs
}

// IsBodyMap gets whether the body is a map[string]interface{}.
func (r *Response) IsBodyMap() bool {
	var obj map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &obj); err != nil {
		return false
	}
	return true
}

// IsBodyMapSlice gets the body is a []map[string]interface{}.
func (r *Response) IsBodyMapSlice() bool {
	var objs []map[string]interface{}
	if err := Unmarshal(r.BodyBytes(), &objs); err != nil {
		return false
	}
	return true
}
