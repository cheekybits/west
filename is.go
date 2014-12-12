package west

import "fmt"

// An A type assertions that can be made against
// a response.
type A struct {
	S int
	B interface{}
}

// Is checks the Response with the specified A.
func (a A) Is(r *Response) error {

	if a.S > 0 {
		if a.S != r.StatusCode {
			return fmt.Errorf("status code %d != %d", r.StatusCode, a.S)
		}
	}

	if a.B != nil {

		switch body := a.B.(type) {
		case string: // literal

			if body != r.BodyString() {
				return fmt.Errorf("body \"%s\" != \"%s\"", r.BodyString(), body)
			}

		default:
			panic(fmt.Sprintf("unsupported body type %T", body))
		}

	}

	return nil

}
