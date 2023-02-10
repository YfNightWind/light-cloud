package main

import (
	"flag"
	"fmt"
	"light-cloud/src/core/internal/config"
	"light-cloud/src/core/internal/handler"
	"light-cloud/src/core/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(
		// 设置 header
		func(header http.Header) {
			header.Set("Access-Control-Allow-Origin", "*")
		},
		nil,
		// 允许跨域地址
		"*",
	))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
