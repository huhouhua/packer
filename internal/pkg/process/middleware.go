package process

type Middleware[TCtx any] func(prev any, ctx *TCtx, next HandlerFunc[TCtx]) error

type IMiddleware[TCtx any] interface {
	Invoke(prev any, ctx *TCtx, next HandlerFunc[TCtx]) error
}
