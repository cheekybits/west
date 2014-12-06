package main

import (
	"os"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/west"
)

// HTTP methods
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

func main() {

	t := new(testing.T)
	w := west.New(t).
		Name("test something").
		Prefix("http://www.domain.com/api/v1/").
		Setup(func() {

	}).
		TearDown(func() {

	})

	if !w.Run(testSomething) {
		os.Exit(1)
	}

}

func testSomething(t west.WT) {
	is := is.New(t)

	// TODO: setup
	defer func() {
		// teardown
	}()

	res := west.R{
		M: GET,
		P: "monkey",
	}.Do(t)

	// TODO: assertions
	is.Equal(res.BodyString, `{"something":true}`)

}
