package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	"sample/gen-go/Sample"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

//定义服务
type Greeter struct {
}

//实现IDL里定义的接口
//SayHello
func (this *Greeter) SayHello(ctx context.Context, u *Sample.User) (r *Sample.Response, err error) {
	strJson, _ := json.Marshal(u)
	return &Sample.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

//GetUser
func (this *Greeter) GetUser(ctx context.Context, uid int32) (r *Sample.Response, err error) {
	return &Sample.Response{ErrCode: 1, ErrMsg: "user not exist."}, nil
}
var defaultCtx = context.Background()

func main() {
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	handler := &Greeter{}
	processor := Sample.NewGreeterProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
	return
}

