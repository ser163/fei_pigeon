# fei_pigeon
  一个go语言编写与飞书机器人接口的程序
## 使用说明

您可以使用以下命令来运行此程序：
* 通过管道发送文本消息：`echo "Hello World" | dpigeon`
* 发送文本消息：`fpigeon -type=text -content="Hello World"`

* 发送链接消息：`fpigeon -type=rich -title="Link Title" -content="Link Content" -url="https://example.com"`


## 配置说明
  将config.yaml.bak 改名为config.yaml
```yaml
webhook: https://oapi.dingtalk.com/robot/send?access_token=YOUR_ACCESS_TOKEN_HERE
secret: YOUR_SECRET_HERE
```
## 编译
Linux编译
```shell
go build -ldflags "-s -w" -o fpingeon main.go
```
windows下交叉编译
```shell
set GOOS=linux
set GOARCH=amd64 
go build -ldflags "-s -w" -o fpingeon main.go
```