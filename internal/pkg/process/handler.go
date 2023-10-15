package process

type HandlerFunc[TCtx any] func(prev any, ctx *TCtx) error
