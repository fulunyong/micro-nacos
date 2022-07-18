package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/" summary:"项目欢迎语"`
}

type IndexRes struct {
	Tips string `json:"tips"`
}
