package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/" method:"get"`
	Name   string `v:"required" dc:"姓名"`
	Age    int    `v:"required" dc:"年龄"`
}
type HelloRes struct{}

type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.Writef("hello %s! your age is %d",
		req.Name, req.Age)
	return
}

func ErrorHandler(r *ghttp.Request) {
	//执行路由回调函数
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.Write("error occurs:", err.Error())
		return
	}
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ErrorHandler)
		group.Bind(new(Hello))
	})
	s.SetPort(8080)
	s.Run()
}
