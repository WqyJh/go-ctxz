package ctxz_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wqyjh/go-ctxz"
)

func TestWithoutCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "key", "value")
	ctx3 := ctxz.WithoutCancel(ctx2)
	cancel()
	assert.ErrorIs(t, context.Canceled, ctx.Err())
	assert.ErrorIs(t, context.Canceled, ctx2.Err())
	assert.NoError(t, ctx3.Err())
	assert.Equal(t, "value", ctx3.Value("key"))

	ctx4 := ctxz.WithoutCancel(ctx3)
	assert.NotEqual(t, ctx2, ctx3)
	assert.Equal(t, ctx3, ctx4)
}

func TestWithNewCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "key", "value")
	ctx3, cancel3 := ctxz.WithNewCancel(ctx2)
	cancel()
	assert.ErrorIs(t, context.Canceled, ctx.Err())
	assert.NoError(t, ctx3.Err())
	cancel3()
	assert.ErrorIs(t, context.Canceled, ctx3.Err())
	assert.Equal(t, "value", ctx3.Value("key"))

	ctx, cancel = context.WithCancel(context.Background())
	ctx2 = context.WithValue(ctx, "key", "value")
	ctx3, cancel3 = ctxz.WithNewCancel(ctx2)
	cancel3()
	assert.ErrorIs(t, context.Canceled, ctx3.Err())
	assert.NoError(t, ctx.Err())
	cancel()
	assert.ErrorIs(t, context.Canceled, ctx.Err())
}
