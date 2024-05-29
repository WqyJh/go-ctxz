package ctxz_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wqyjh/go-ctxz"
)

func TestMaybeTimtout(t *testing.T) {
	ctx, _, _ := ctxz.WithMaybeTimeout(context.Background(), time.Millisecond*50)
	time.Sleep(time.Millisecond * 100)
	assert.ErrorIs(t, ctx.Err(), context.Canceled)
	t.Logf("test error: %v", ctx.Err())

	ctx, cancel, _ := ctxz.WithMaybeTimeout(context.Background(), time.Millisecond*50)
	cancel()
	assert.ErrorIs(t, ctx.Err(), context.Canceled)

	ctx, cancel, notify := ctxz.WithMaybeTimeout(context.Background(), time.Millisecond*50)
	notify()
	notify() // support reentrant
	time.Sleep(time.Millisecond * 100)
	assert.NoError(t, ctx.Err())
	cancel()
	assert.ErrorIs(t, ctx.Err(), context.Canceled)
}
