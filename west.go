package west

import "github.com/cheekybits/is"

type W interface {
	Run(...WFunc) bool
	Prefix(string) W
	Name(string) W
}
type T interface {
	is.T
}
type WT interface {
	T
}
type WFunc func(WT)

func New(t T) W {
	return &w{t: t}
}

type w struct {
	t is.T
}

func (w *w) Run(fns ...WFunc) bool {
	return true
}
func (w *w) Prefix(p string) W {

	return w
}
func (w *w) Name(n string) W {

	return w
}
