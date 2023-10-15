package process

var _ IBuilder[any] = &Builder[any]{}

type IBuilder[TCtx any] interface {
	UseWhen(middleware IMiddleware[TCtx]) IBuilder[TCtx]
	Use(middleware Middleware[TCtx]) IBuilder[TCtx]
	Build() HandlerFunc[TCtx]
}
type Builder[TCtx any] struct {
	mdls             []middlewareInvoke[TCtx]
	completedHandler HandlerFunc[TCtx]
}

func NewBuilder[TCtx any]() IBuilder[TCtx] {
	return &Builder[TCtx]{
		mdls: []middlewareInvoke[TCtx]{},
		completedHandler: func(prev any, ctx *TCtx) error {
			return nil
		}}
}

type middlewareInvoke[TCtx any] func(handlerFunc HandlerFunc[TCtx]) HandlerFunc[TCtx]

func (b *Builder[TCtx]) UseWhen(middleware IMiddleware[TCtx]) IBuilder[TCtx] {
	return b.Use(middleware.Invoke)
}

func (b *Builder[TCtx]) Use(middleware Middleware[TCtx]) IBuilder[TCtx] {
	return b.add(func(next HandlerFunc[TCtx]) HandlerFunc[TCtx] {
		return func(prev any, ctx *TCtx) error {
			return middleware(prev, ctx, next)
		}
	})
}

func (b *Builder[TCtx]) add(invoke middlewareInvoke[TCtx]) IBuilder[TCtx] {
	b.mdls = append(b.mdls, invoke)
	return b
}
func (b *Builder[TCtx]) Build() HandlerFunc[TCtx] {
	root := b.completedHandler
	for i := len(b.mdls) - 1; i >= 0; i-- {
		root = b.mdls[i](root)
	}
	return root
}
