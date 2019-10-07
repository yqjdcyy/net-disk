package config

import (
)

// GatewayConfig  网关服务的配置
type GatewayConfig struct {
	c *ConfigFile

	// Port 服务启动端口
	Port int

	// Dir 基础目录地址
	Dir string

	// Suffix 缩略图后缀
	Suffix string
}

func (g *GatewayConfig) init(c *ConfigFile) {
	g.c = c

	g.Port = c.GetIntDefault("setting", "port", 7000)
	
	g.Dir = c.GetStringDefault("resource", "dir", "/data/cdn/dev/img")
	g.Suffix = c.GetStringDefault("resource", "suffix", "")
}
