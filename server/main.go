package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sample/gen-go/Sample"
	"sample/servers"
	"time"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

//定义服务
type Greeter struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	Addr     string
	Port     int
}

//实现IDL里定义的接口
//SayHello
func (this *Greeter) SayHello(ctx context.Context, u *Sample.User) (r *Sample.Response, err error) {
	fmt.Println("say helloword!")
	strJson, _ := json.Marshal(u)
	return &Sample.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

//GetUser
func (this *Greeter) GetUser(ctx context.Context, uid int32) (r *Sample.Response, err error) {
	fmt.Println("say GetUser!")
	return &Sample.Response{ErrCode: 1, ErrMsg: "user not exist."}, nil
}

var defaultCtx = context.Background()

func main() {
	srv, err := servers.NewServers()
	if err != nil {
		panic("create server failed")
		return
	}
	handler := &Greeter{}
	processor := Sample.NewGreeterProcessor(handler)
	servers.RegisterService("trpc.ten_video_live.live_management_log.live_management_log", processor)
	err = srv.Server()
	if err != nil {
		panic(err)
	}
	select {} // 阻塞
	return
}
