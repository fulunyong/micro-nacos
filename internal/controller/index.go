package controller

import (
	"context"
	v1 "micro-nacos/api/v1"
)

var Index = cUser{}

type cUser struct{}

func (c *cUser) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	return &v1.IndexRes{Tips: "项目启动成功！"}, err
}
