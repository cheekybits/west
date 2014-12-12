package west_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/west"
)

func TestResponseIs(t *testing.T) {
	is := is.New(t)

	res := &west.Response{Response: &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`string`)),
	}}
	is.NoErr(west.A{S: 200, B: "string"}.Is(res))
	is.Equal("status code 200 != 201", west.A{S: 201}.Is(res).Error())
	is.Equal("body \"string\" != \"strings\"", west.A{B: "strings"}.Is(res).Error())

}
