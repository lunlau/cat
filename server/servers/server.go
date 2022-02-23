package servers

import (
	"fmt"
	"sample/filters"
	"time"
)

var serverGlobal *ServerFactory

func GetGlobalServer() *ServerFactory {
	if serverGlobal == nil {
		panic("server global is nil")
		return nil
	}
	return serverGlobal
}

func NewServers(opt ...OptionFuncType) (*ServerFactory, error) {
	// new 服务
	serverGlobal = &ServerFactory{
		SrvGlobalConfig: &Config{},
	}
	// 初始化配置
	InitConfig(serverGlobal.SrvGlobalConfig)
	// 遍历producer服务配置
	for _, v := range serverGlobal.SrvGlobalConfig.Server.Service {
		service := newService(v)
		// 注册service
		serviceName2ProxyMap[v.Name] = service
	}
	// 遍历consumer客户端配置
	// 加载插件
	err := serverGlobal.SrvGlobalConfig.lugins.Run()
	if err != nil {
		return serverGlobal, err
	}
	return serverGlobal, nil
}
func newService(service *ProducerSrvConfig) *ServiceProxy {
	if service == nil {
		panic("service is nil")
	}
	// 配置文件注册
	return &ServiceProxy{
		cfg: service,
	}
}

// Server 启动服务
func (s *ServerFactory) Server() error {
	// 启动配置后的service
	for sn, s := range serviceName2ProxyMap {
		go func() {
			err := s.Server()
			if err != nil {
				panic(fmt.Errorf("service name %v run failed %+v", sn, err))
				return
			}
		}()

	}
	return nil
}

type ServerFactory struct {
	SrvGlobalConfig *Config
}

// ServerOptions 服务参数
type ServerOptions struct {
	Address    string        // 监听地址 ip:port
	Timeout    time.Duration // 请求最长处理时
	FilterList filters.Chain // 链式拦截器
}

// Option 服务启动参数工具函数
type OptionFuncType func(*ServerOptions)

func WithAddress(addr string) OptionFuncType {
	return func(options *ServerOptions) {
		options.Address = addr
	}
}

func WithTimeout(t time.Duration) OptionFuncType {
	return func(options *ServerOptions) {
		options.Timeout = t
	}
}

func WithFilters(filterList filters.Chain) OptionFuncType {
	return func(options *ServerOptions) {
		options.FilterList = filterList
	}
}
