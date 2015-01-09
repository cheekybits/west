package west

import (
	"reflect"
	"strings"

	"fmt"
)

// An A type assertions that can be made against
// a response.
type A struct {
	// S is the expected status code.
	S int
	// B is the expected body object.
	B interface{}
	// H is a map of expected (response) headers from the server.
	H map[string]string
}

// Is checks the Respons with the specified A.
// Returns nil if everything is OK.
func (r *Response) Is(a A) error {
	return a.Is(r)
}

// Is checks the Response with the specified A.
// Returns nil if everything is OK.
func (a A) Is(r *Response) error {

	// check status code
	if a.S > 0 {
		if a.S != r.StatusCode {
			return fmt.Errorf("status code %d != %d", r.StatusCode, a.S)
		}
	}

	// check headers
	if a.H != nil {
		for k, v := range a.H {
			actual := r.Header.Get(k)
			if !strings.Contains(actual, v) {
				return fmt.Errorf("header '%s' doesn't contain '%s': %s", k, v, actual)
			}
		}
	}

	// check body
	if a.B != nil {

		switch body := a.B.(type) {
		case string: // literal
			if body != r.BodyString() {
				return fmt.Errorf("body string \"%s\" != \"%s\"", r.BodyString(), body)
			}
		case []byte:
			if string(body) != string(r.BodyBytes()) {
				return fmt.Errorf("body bytes \"%s\" != \"%s\"", string(r.BodyString()), string(body))
			}
		case map[string]interface{}:
			if !reflect.DeepEqual(body, r.BodyMap()) {
				return fmt.Errorf("body %v != %v", body, r.BodyMap())
			}
		default:
			expBytes, err := Marshal(body)
			if err != nil {
				return err
			}
			actBytes, err := Marshal(r.BodyObj())
			if err != nil {
				return err
			}
			expected := string(expBytes)
			actual := string(actBytes)
			if actual != expected {
				return fmt.Errorf("body \"%s\" != \"%s\"", actual, expected)
			}
		}

	}

	return nil

}
