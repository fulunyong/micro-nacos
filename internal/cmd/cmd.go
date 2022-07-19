package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"micro-nacos/internal/controller"
	"micro-nacos/nacos"
	"net"
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
						"/echo": controller.Index.Echo,
						"/":     controller.Index.Index,
					})
				})
			})
			// Custom enhance API document.
			// Just run the server.
			port := s.GetListenedPort()
			if 0 == port {
				//如果没有配置端口  随机获取一个可用端口
				port = GetPort()
				s.SetPort(port)
			}
			nacos.RegisterServe(port)
			fmt.Printf("当前服务端口:%d\n", port)
			s.Run()
			return nil
		},
	}
)

func GetPort() int {
	address, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
		return 0
	}
	listen, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}
