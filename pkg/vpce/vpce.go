package vpce

import "errors"

// Manager interface
type Manager interface {
	JobTracer(ctx *Context, fn HandlerFunc) JobTracer
	Detail(ctx *Context, fn HandlerFunc) (*Spec, error)
}

type HandlerFunc func(ctx *Context) (interface{}, error)

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}

func (m *manager) JobTracer(ctx *Context, fn HandlerFunc) JobTracer {
	arg, err := fn(ctx)
	if err != nil {
		return &jobTracer{err: err}
	}

	return &jobTracer{arg: arg}
}

func (m *manager) Detail(ctx *Context, fn HandlerFunc) (*Spec, error) {
	data, err := fn(ctx)
	if err != nil {
		return nil, err
	}

	spec, ok := data.(*Spec)
	if !ok {
		return nil, errors.New("invalid data type")
	}

	return spec, nil
}

// JobTracer interface
type JobTracer interface {
	Result(fn JobFunc) (*JobDetail, error)
}

type JobFunc func(arg interface{}) (*JobDetail, error)

type jobTracer struct {
	arg interface{}

	err error
}

func isZero(args interface{}) bool {
	if args == nil {
		return true
	}

	val, ok := args.(int)
	if ok {
		return val == 0
	}

	str, ok := args.(string)
	if ok {
		return len(str) == 0
	}

	return true
}

func (j *jobTracer) Result(fn JobFunc) (*JobDetail, error) {
	if isZero(j.arg) {
		return nil, j.err
	}

	return fn(j.arg)
}

// JobDetail struct
type JobDetail struct {
}
