package config

import(
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	
	"gopkg.in/yaml.v2"
)


// ClientConf 客户端配置
var ClientConf *Confige
// Path 配置文件目录
var Path string
// Quit 退出标志
var Quit chan bool

func init() {
	flag.StringVar(&Path,"p", "../properties/conf.yaml", "通用系统配置文件目录")
	Quit = make(chan bool)
}

// Init 使用默认配置进行初始化
func Init() {

	// init
	ClientConf = &Confige{
		LogPath: "log/client.log",
	}

	// read
	r, err := os.Open(Path)
	if nil != err {
		fmt.Printf("fail to read file[%s]: %s", Path, err.Error())
		return
	}
	bs, err := ioutil.ReadAll(r)
	if nil != err {
		fmt.Printf("fail to read file[%s].all: %s", Path, err.Error())
		return
	}

	// parse
	err = yaml.Unmarshal(bs, ClientConf)
	if nil != err {
		fmt.Printf("fail to unmarshal data[%s] to Confige: %s", string(bs), err.Error())
		return
	}

	// check
	// fmt.Println(ClientConf.URL)
	// fmt.Println(ClientConf.LogPath)
	// fmt.Println(ClientConf.Path)
	// fmt.Println(ClientConf.Duration)
}

// Confige 配置
type Confige struct {
	// Url 请求地址
	URL string `yaml:"url,omitempty"`
	// LogPath 日志目录
	LogPath string `yaml:"log_path,omitempty"`
	// Path 配置路径
	Path []string `yaml:"path,omitempty"`
	// Duration 定时任务间隔，单位为分钟
	Duration int `yaml:"duration"`
	// Type 兼容类型
	Type []string `yaml:"type,omitempty"`
	// Delete
	Delete bool `yaml:"delete"`
}
