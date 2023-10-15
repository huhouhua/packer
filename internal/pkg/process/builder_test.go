package process

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_builder(t *testing.T) {

	build := NewBuilder[Context]().Use(func(prev any, ctx *Context, next HandlerFunc[Context]) error {
		t.Log(prev)
		_ = next("me1", ctx)
		return nil
	}).Use(func(prev any, ctx *Context, next HandlerFunc[Context]) error {
		t.Log(prev)
		_ = next("me2", ctx)
		return nil
	}).Use(func(prev any, ctx *Context, next HandlerFunc[Context]) error {
		t.Log(prev)
		_ = next("me3", ctx)
		return nil
	}).Build()

	ctx := &Context{}
	err := build("m0", ctx)
	require.NoError(t, err)
}
