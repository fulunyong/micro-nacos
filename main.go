package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"micro-nacos/internal/cmd"
)

func main() {
	fmt.Println("服务启动 Hello, World!")
	cmd.Main.Run(gctx.New())
}
