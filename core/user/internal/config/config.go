package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL struct {
		DataSourceName string
	}
	Redis struct {
		Address  string
		Password string
	}
}
