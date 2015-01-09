package west

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// R is a request.
type R struct {
	// M is the HTTP method.
	M string
	// P is the path.
	P string
	// B is the body. Can be string, []byte, io.Reader.
	// Anything else will be marshalled.
	B interface{}
	// H represents the request headers to be sent to the
	// server.
	H map[string]string
	// Client is the http.Client to use when making
	// requests.
	Client *http.Client
}

// MustDo makes the request against the specified httptest.Server and
// returns a Response, or panics if an error occurs.
func (r R) MustDo(s *httptest.Server) *Response {
	res, err := r.Do(s)
	if err != nil {
		panic("MustDo failed: " + err.Error())
	}
	return res
}

// Do makes the request against the specified httptest.Server and
// returns a Response, or an error.
func (r R) Do(s *httptest.Server) (*Response, error) {

	// setup defaults
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	// make url
	u, err := url.Parse(s.URL + "/" + strings.TrimPrefix(r.P, "/"))
	if err != nil {
		return nil, err
	}

	// make request
	var req *http.Request
	var bodyReader io.Reader

	if r.B != nil {
		switch body := r.B.(type) {
		case string:
			bodyReader = strings.NewReader(body)
		case []byte:
			bodyReader = bytes.NewReader(body)
		case io.Reader:
			bodyReader = body
		default:
			b, err := Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(b)
		}
	}

	req, err = http.NewRequest(r.M, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// set headers
	if r.H != nil {
		for k, v := range r.H {
			req.Header.Set(k, v)
		}
	}

	// make request
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{Response: resp}, nil
}
