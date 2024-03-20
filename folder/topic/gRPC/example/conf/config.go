package conf

import (
	beeConfig "github.com/astaxie/beego/config"
)

var Configer *config

type config struct {
	BeeConfiger beeConfig.Configer
	RpcConf     RpcConfig
	RedisConf   RedisConfig
}

func NewConfiger(filename string) {

	Configer = new(config)
	var err error

	Configer.BeeConfiger, err = beeConfig.NewConfig("ini", filename)
	if err != nil {
		panic("读取配置文件出错")
	}

	Configer.load()
}

func (c *config) load() {

	var err error

	c.RpcConf.RpcPort = c.BeeConfiger.String("system::rpc_port")
	c.RpcConf.LogPath = c.BeeConfiger.String("log::path")
	c.RpcConf.LogLevel = c.BeeConfiger.String("log::level")
	if c.RpcConf.LogLevel != "debug" && c.RpcConf.LogLevel != "info" && c.RpcConf.LogLevel != "error" {
		panic("log_level 配置错误")
	}

	if c.BeeConfiger.String("redis::type") != "" {
		c.RedisConf.Type, err = c.BeeConfiger.Int("redis::type")
		if err != nil {
			panic("读取 redis::type 配置出错")
		}
		if c.RedisConf.Type != 1 && c.RedisConf.Type != 2 {
			panic("读取 redis::type 配置出错，只能为 1、2")
		}

		switch c.RedisConf.Type {
		case 2: // 集群
			c.RedisConf.Cluster.Addrs = c.BeeConfiger.Strings("redis::addrs")
		default: // 单点
			c.RedisConf.Single.Host = c.BeeConfiger.String("redis::host")
			c.RedisConf.Single.Port = c.BeeConfiger.String("redis::port")
			c.RedisConf.Single.DB, err = c.BeeConfiger.Int("redis::db")
			if err != nil {
				panic("读取 redis::db 配置出错")
			}
		}

		c.RedisConf.Password = c.BeeConfiger.String("redis::password")

		if c.BeeConfiger.String("redis::pool_size") != "" {
			c.RedisConf.PoolSize, err = c.BeeConfiger.Int("redis::pool_size")
			if err != nil {
				panic("读取 redis::pool_size 配置出错")
			}
		}
	}

}
