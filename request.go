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
	// Client is the http.Client to use when making
	// requests.
	Client *http.Client
}

// Do makes the request against the specified httptest.Server and
// returns a Res response, or an error.
func (r R) Do(s *httptest.Server) (*Resp, error) {

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

	if r.B != nil {
		var bodyReader io.Reader
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
		req, err = http.NewRequest(r.M, u.String(), bodyReader)
		if err != nil {
			return nil, err
		}
	} else {
		// no body
		req, err = http.NewRequest(r.M, u.String(), nil)
		if err != nil {
			return nil, err
		}
	}

	// make request
	response, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Resp{Response: response}, nil
}