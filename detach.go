package ctxz

import (
	"context"
	"time"
)

// WithoutCancel returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling.
// https://github.com/golang/tools/blob/master/internal/xcontext/xcontext.go
func WithoutCancel(ctx context.Context) context.Context {
	_, ok := ctx.(detachedContext)
	if ok {
		// no need to detach twice
		return ctx
	}
	return detachedContext{ctx}
}

type detachedContext struct{ parent context.Context }

func (v detachedContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (v detachedContext) Done() <-chan struct{}             { return nil }
func (v detachedContext) Err() error                        { return nil }
func (v detachedContext) Value(key interface{}) interface{} { return v.parent.Value(key) }

// WithNewCancel returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling and provides a new cancel function.
func WithNewCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx = WithoutCancel(ctx)
	return context.WithCancel(ctx)
}
