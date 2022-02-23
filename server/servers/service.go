package servers

import (
	"crypto/tls"
	"fmt"
	"sample/components/log"

	"github.com/apache/thrift/lib/go/thrift"
)

func (s *ServiceProxy) Server() error {
	if s.processor == nil {
		log.Info("processor is empty")
		return nil
	}
	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	fmt.Println(cfg)
	addr := fmt.Sprintf("%s:%d", s.cfg.Ip, s.cfg.Port)
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		panic(err)
	}
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		s.processor,
		transport,
		transportFactory,
		protocolFactory,
	)
	err = server.Serve()
	if err != nil {
		fmt.Println("error running server:", err)
		return err
	}
	s.server = server
	return nil
}

var serviceName2ProxyMap = map[string]*ServiceProxy{}

// 处理配置文件service
func RegisterService(serviceName string, processor thrift.TProcessor) {
	v, ok := serviceName2ProxyMap[serviceName]
	if !ok {
		panic("service name is invalid ")
	}
	v.processor = processor

}

type ServiceProxy struct {
	server    *thrift.TSimpleServer
	processor thrift.TProcessor
	cfg       *ProducerSrvConfig
}
