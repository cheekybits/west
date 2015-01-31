package west_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/west"
)

func TestResponseBodyBytes(t *testing.T) {
	is := is.New(t)
	res := &west.Response{Response: &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`bytes`)),
	}}
	is.Equal([]byte(`bytes`), res.BodyBytes())
}

func TestResponseBodyString(t *testing.T) {
	is := is.New(t)
	res := &west.Response{Response: &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`string`)),
	}}
	is.Equal(`string`, res.BodyString())
}

func TestResponseBodyMap(t *testing.T) {
	is := is.New(t)
	res := &west.Response{Response: &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"name":"cheekybits"}`)),
	}}
	is.Equal(`cheekybits`, res.BodyMap()["name"])
}

func TestResponseBodyUnmarshal(t *testing.T) {
	is := is.New(t)
	res := &west.Response{Response: &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"name":"cheekybits"}`)),
	}}
	o := struct {
		Name string
	}{}
	res.UnmarshalBody(&o)
	is.Equal(`cheekybits`, o.Name)
}

func TestResponseBodyMapSlice(t *testing.T) {
	is := is.New(t)
	res := &west.Response{Response: &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`[{"name":"cheekybits"},{"name":"verycheekybits"}]`)),
	}}
	is.Equal(`cheekybits`, res.BodyMapSlice()[0]["name"])
	is.Equal(`verycheekybits`, res.BodyMapSlice()[1]["name"])
}
