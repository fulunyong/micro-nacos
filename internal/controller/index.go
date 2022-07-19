package controller

import (
	"context"
	"fmt"
	v1 "micro-nacos/api/v1"
	"micro-nacos/nacos"
)

var Index = cUser{}

type cUser struct{}

func (c *cUser) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	fmt.Println("请求项目消息")
	return &v1.IndexRes{Tips: "项目启动成功！"}, err
}
func (c *cUser) Echo(ctx context.Context, req *v1.EchoReq) (res *v1.IndexRes, err error) {
	fmt.Println("欢迎访问数据")
	return &v1.IndexRes{Tips: "im form :" + nacos.MustGetIntranetIp()}, err
}
