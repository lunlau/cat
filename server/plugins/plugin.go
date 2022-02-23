package plugins

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/dullgiulio/pingo"
	yaml "gopkg.in/yaml.v3"
)

func test() {
	pingo.Register(&struct{}{})
	pingo.Run()
}

// Register a new object this plugin exports. The object must be
// an exported symbol and obey all rules an object in the standard
// "rpc" module has to obey.
//
// Register will panic if called after Run.
func Register(pluginName string, obj interface{}) {
	if defaultPluginFactory.running {
		panic("Do not call Register after Run")
	}
	defaultPluginFactory.register(pluginName, obj)
}

var defaultPluginFactory = newPluginFactory()

// PluginsConfig type name  pluginconfig
type PluginsConfig map[string]map[string]yaml.Node

func (cfg *PluginsConfig) Run() error {
	for pluginType, v := range *cfg {
		fmt.Println("plugin type ", pluginType)
		// 找到插件
		// 遍历插件名称
		for pluginID, pluginNode := range v {
			pluginObj := getPluginByName(pluginID)
			pluginObjImp, ok := pluginObj.(PluginParent)
			if !ok {
				return fmt.Errorf("plugin is invalid")
			}
			yamlDecoder := &YamlNodeDecoder{
				Node: &pluginNode,
			}
			pluginObjImp.Run(pluginID, yamlDecoder)
		}
	}
	return nil
}

// 插件注册
func (r *pluginFactory) register(pluginName string, obj interface{}) {
	element := reflect.TypeOf(obj).Elem()
	r.objs = append(r.objs, element.Name())
	r.pluginsMap.Store(pluginName, obj)
}

func getPluginByName(pluginId string) interface{} {
	v, ok := defaultPluginFactory.pluginsMap.Load(pluginId)
	if !ok {
		return nil
	}
	return v
}

// 插件启动
func (r *pluginFactory) run() error {

	return nil
}

type pluginFactory struct {
	secret     string
	objs       []string
	conf       *PluginsConfig
	running    bool
	pluginsMap sync.Map // map[string]*service
}

func makeConfig() *PluginsConfig {
	return &PluginsConfig{}
}

func newPluginFactory() *pluginFactory {
	rand.Seed(time.Now().UTC().UnixNano())
	r := &pluginFactory{
		secret:     randstr(64),
		objs:       make([]string, 0),
		conf:       makeConfig(), // conf remains fixed after this point
		pluginsMap: sync.Map{},
	}
	return r
}

var _letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-")

func randstr(n int) string {
	b := make([]rune, n)
	l := len(_letters)
	for i := range b {
		b[i] = _letters[rand.Intn(l)]
	}
	return string(b)
}

// plugin 1、注册 2、初始化 3、工厂获取对象
// 1、定义plugin， run函数，包括启动函数
type PluginParent interface {
	PluginParentType() string
	Run(pluginID string, decoder Decoder) error
}

// Decoder 节点配置解析器。
type Decoder interface {
	Decode(cfg interface{}) error // 输入参数为自定义的配置数据结构
}

// YamlNodeDecoder yaml节点配置解析器。
type YamlNodeDecoder struct {
	Node *yaml.Node
}

// Decode 解析yaml node配置。
func (d *YamlNodeDecoder) Decode(cfg interface{}) error {
	if d.Node == nil {
		return errors.New("yaml node empty")
	}
	return d.Node.Decode(cfg)
}
