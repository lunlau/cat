package servers

import (
	"fmt"
	"io/ioutil"
	"log"
	"sample/plugins"

	yaml "gopkg.in/yaml.v3"
)

var cfgPathDefault = "./server.yaml"

func InitConfig(srvGlobalConfig *Config) error {
	if srvGlobalConfig == nil {
		return fmt.Errorf("server cfg is nil")
	}
	var cfgPath string // 从环境变量取值
	if cfgPath != "" {
		cfgPathDefault = cfgPath
	}
	yamlFile, err := ioutil.ReadFile(cfgPathDefault)
	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, srvGlobalConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	log.Println("conf", srvGlobalConfig)
	return nil
}

// Config trpc配置实现，分四大块：全局配置global，服务端配置server，客户端配置client，插件配置plugins
type Config struct {
	Global struct {
		EnvName       string `yaml:"env_name"`
		ContainerName string `yaml:"container_name"`
		LocalIP       string `yaml:"local_ip"`
	}
	Server struct {
		Filter  []string             // 针对所有service的拦截器
		Service []*ProducerSrvConfig // 单个service服务的配置
	}
	ConsumerCfg ConsumerConfigs
	lugins      plugins.PluginsConfig
}

type ConsumerConfigs struct {
	Filter         []string // 针对所有后端的拦截器
	Namespace      string   // 针对所有后端的namespace
	Timeout        int
	Discovery      string
	Loadbalance    string
	Circuitbreaker string
	Service        []*ConsumerSrvConfig // 单个后端消费者请求的配置
}

type ConsumerSrvConfig struct {
	EnvName  string `yaml:"env_name"` // 设置下游服务多环境的环境名
	Target   string // url配置
	Password string
	Timeout  int // 单位 ms
	Filter   []string
}
type ProducerSrvConfig struct {
	Name     string // 服务名
	EnvName  string `yaml:"env_name"` // 设置下游服务多环境的环境名
	Ip       string // url配置
	Port     int
	Password string
	Timeout  int // 单位 ms
	Filter   []string
}

// Run producer启动
func (s *ProducerSrvConfig) Setup() {

}
