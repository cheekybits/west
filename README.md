# west

Package west provides helpers for testing HTTP endpoints with the httptest package.

### How it works

```
func TestSomething(t *testing.T) {
  is := is.New(t) // use whichever framework you like

  // use httptest to make a test server for YourHandler
  s := httptest.NewServer(new(YourHandler))
  defer s.Close() // always do this right away

  // use west.R to make a request, and call Do
  res, err := west.R{
    M: "GET", P: "/path",
    B: "body",
  }.Do(s)

  // assert that no error occurred
  is.NoErr(err)

  // check some things aobut the response
  is.Equal(http.StatusOK, res.StatusCode)
  is.Equal("application/json", res.Header.Get("Content-Type"))

  // assuming the body was a JSON object, get it and make
  // some assertions about it
  obj := res.BodyMap()
  is.Equal("cheekybits", obj["name"])

}
```

