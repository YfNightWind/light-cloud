# Light-Cloud轻量级网盘系统(Go)
### 目前正在升级到微服务架构，详细可见[microservice分支](https://github.com/YfNightWind/light-cloud/tree/microservice)
### 运行前您需要知道的
- 进入到```core/model/init.go```中，修改对应的参数以保证能正常连接您的数据库  
  `goctl api go -api core.api -dir . -style go_zero`  
  运行`go run core.go -f etc/core-api.yaml`
