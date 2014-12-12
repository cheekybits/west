package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/west"
)

type EchoServer struct{}

var _ http.Handler = (*EchoServer)(nil)

func (t *EchoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200
	statusStr := r.URL.Query().Get("status")
	if len(statusStr) > 0 {
		statusInt64, err := strconv.ParseInt(statusStr, 10, 32)
		if err != nil {
			panic(err)
		}
		status = int(statusInt64)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	b, _ := ioutil.ReadAll(r.Body)
	data := map[string]interface{}{
		"path": r.URL.Path,
		"body": string(b),
	}
	json.NewEncoder(w).Encode(data)
}

func TestCRUD(t *testing.T) {
	is := is.New(t)

	// use httptest to make a test server for YourHandler
	s := httptest.NewServer(new(EchoServer))
	defer s.Close() // always do this right away

	is.NoErr(west.R{
		P: "/something",
	}.MustDo(s).Is(west.A{
		S: 200,
		B: map[string]interface{}{"body": "", "path": "/something"},
	}))

	is.NoErr(west.R{
		P: "/something",
		B: `literal body`,
	}.MustDo(s).Is(west.A{
		S: 200,
		B: map[string]interface{}{"body": "literal body", "path": "/something"},
	}))

}
