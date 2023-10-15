package process

type RegisterFunc[TCtx any] func(ctx *TCtx) error
