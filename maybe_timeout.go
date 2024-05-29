package ctxz

import (
	"context"
	"time"
)

type CancelFunc = context.CancelFunc
type NotifyFunc func()

func WithMaybeTimeout(ctx context.Context, timeout time.Duration) (context.Context, CancelFunc, NotifyFunc) {
	var notify = make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-notify:
			return
		case <-ctx.Done():
			return
		case <-time.After(timeout):
			cancel()
		}
	}()
	return ctx,
		func() {
			cancel()
		},
		func() {
			select {
			case notify <- struct{}{}:
			default:
			}
		}
}
