# HelloGo
A simple go project created by beego.  

### [Go 离线版教程](https://tour.go-zh.org/welcome/3)
```command line
go get -u github.com/Go-zh/tour tour
```

### 日志

##### 支持多个配置文件
- 使用 include 方式，引用多个配置文件
> 如:在 app.conf 中添加一行 `include "logs.conf"`

##### 支持环境变量配置
- 配置文件解析支持从环境变量中获取配置项，配置项格式：\${环境变量}
> 如 : `log_path="${GOLOGPATH||logs}"` 
> 如果有配置环境变量 `GOLOGPATH` 则优先使用该环境变量值。如果不存在或者为空则使用 `logs`作为默认值