package filters

import "context"

// Chain 链式过滤器
type Chain []Filter

// ChainDefault 默认的拦截器链
var ChainDefault = Chain{}

// Filter 过滤器（拦截器），根据dispatch处理流程进行上下文拦截处理
type Filter func(ctx context.Context, req interface{}, rsp interface{}, f HandleFilter) (err error)

// HandleFunc 过滤器（拦截器）函数接口
type HandleFilter func(ctx context.Context, req interface{}, rsp interface{}) (err error)

// Handle 链式过滤器递归处理流程
func (fc Chain) Handle(ctx context.Context, req, rsp interface{}, f HandleFilter) (err error) {
	if len(fc) == 0 {
		return f(ctx, req, rsp)
	}
	return fc[0](ctx, req, rsp, func(ctx context.Context, req interface{}, rsp interface{}) (err error) {
		return fc[1:].Handle(ctx, req, rsp, f)
	})
}
