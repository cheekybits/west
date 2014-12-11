package west_test

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
		"path":  r.URL.Path,
		"query": r.URL.Query(),
		"body":  string(b),
	}
	json.NewEncoder(w).Encode(data)
}

func TestWest(t *testing.T) {
	is := is.New(t)
	s := httptest.NewServer(new(EchoServer))
	defer s.Close()
	res, err := west.R{
		P: "something?name=cheekybits",
		B: "Hello world",
	}.Do(s)
	is.NoErr(err)
	is.Equal(res.StatusCode, http.StatusOK)
	is.Equal(res.BodyString(), `{"body":"Hello world","path":"/something","query":{"name":["cheekybits"]}}`+"\n")
	is.Equal(res.Header.Get("Content-Type"), "application/json")
}
