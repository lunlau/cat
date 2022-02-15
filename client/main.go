package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"os"
	"sample/gen-go/Sample"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport, err := transportFactory.GetTransport(transport)
	client := Sample.NewGreeterClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	res, err := client.GetUser(context.Background(), 1)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("response:", res)
	fmt.Println("well done")
}