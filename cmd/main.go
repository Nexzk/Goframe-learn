package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2"
)

type HelloReq struct {
	name string // 姓名
	age  int    // 年龄
}

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef("hello %s!,your age is %d", r.Get("name", "unknow").String(), r.Get("age", 18).Int())
	})
	s.SetPort(8080)
	s.Run()
	fmt.Println("Hello Gofraemwork!:", gf.VERSION)
}
