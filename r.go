package west

type R struct {
	M string
	P string
}

type Res struct {
	BodyString string
}

func (r R) Do(t WT) *Res {

	// make the request, and return
	// the response

	return &Res{}

}
