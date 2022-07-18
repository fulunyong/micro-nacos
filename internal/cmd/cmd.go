package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"micro-nacos/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple go frame demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					//service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				// Register route handlers.
				//group.Bind(
				//	controller.index,
				//)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/": controller.Index.Index,
					})
				})
			})
			// Custom enhance API document.
			// Just run the server.
			s.Run()
			return nil
		},
	}
)
