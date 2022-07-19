package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"micro-nacos/internal/cmd"
	"micro-nacos/nacos"
)

func main() {
	fmt.Println("服务启动 Hello, World!")
	nacos.RegisterServe()
	cmd.Main.Run(gctx.New())
}
