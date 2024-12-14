package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	Name string // 姓名
	Age  int    // 年龄
}

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		var req HelloReq
		if err := r.Parse(&req); err != nil {
			r.Response.Write(err.Error())
			return
		}
		if req.Name == "" {
			r.Response.Write("name should not be empty")
			return
		}
		if req.Age <= 0 {
			r.Response.Write("invalid age value")
			return
		}
		r.Response.Writef("hello %s! your age is %d", req.Name, req.Age)
	})
	s.SetPort(8080)
	s.Run()
}
