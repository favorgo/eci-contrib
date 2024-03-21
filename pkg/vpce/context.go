package vpce

import "context"

type Context struct {
	ctx context.Context

	Spec *Spec
}

type Spec struct {
}

func NewContext() *Context {
	return &Context{}
}
