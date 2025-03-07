# 版本 V1.0.0

# go-initial

## 目录
- [项目说明](#项目说明)
- [安装依赖](#安装依赖)
- [使用方式](#使用方式)
- [主要功能](#主要功能)

## 项目说明
GO项目初始化文件
#### 目录结构
```
.
├── cmd/                        
│   ├── api/   API服务入口,主服务向外提供接口
│   │   ├── config.yaml         配置文件
│   │   ├── locales             国际化语言文件
│   │   ├── log                 日志文件
│   │   ├── static/             静态文件图片资源等(即将废弃⚠️)
│   │   ├── main_test.go        测试入口文件
│   │   └── main.go             服务入口文件
├── docker/                  docker镜像打包配置文件(DockerFile)
├── internal/                服务内部逻辑处理
│   ├── api/                    外部接口
│   ├── config/                 应用配置
│   ├── db/                     数据库
│   ├── handler/                中间件 
│   ├── scheduler/              调度任务
│   └── services/               业务逻辑
├── pkg/                     公共库   
├── test/                    单元测试 
├── .env                     环境变量
├── tob_group.yml            项目DockerCompose启动配置文件             
├── go.mod
├── go.sum           
├── vendor/                  第三方依赖                
└── README.md                项目说明
```
## 安装依赖
```
go mod tidy

protobuf 格式生成
https://github.com/protocolbuffers/protobuf/releases 安装对应版本 添加到环境变量使用
vi ~/.zshrc or ~/.bash_profile
source ~/.bash_profile
source ~/.zshrc
下载项目 google.golang.org/protobuf/cmd/protoc-gen-go
进入 cmd/protoc-gen-go 目录下运行 go build
生成可执行二进制文件 丢到GOPATH/bin目录下
将 GOPATH/bin 添加到你的 PATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
protoc-gen-go --version 测试是否安装完成

protoc --version 测试是否正确安装
protoc --go_out=. device.proto #生成 Protocol协议对应的go消息结构文件
```
## Docker部署服务
```
docker compose -f tob_group.yml up -d 启动所有容器服务
```

#### 使用 Kafka 消息队列实现读写分离

1. **Kafka 简介**：Kafka 是由 Apache 软件基金会开发的一个开源流处理平台，由 Scala 和 Java 编写。该项目的目标是为处理实时数据提供一个统一、高吞吐、低延迟的平台。其持久化层本质上是一个“按照分布式事务日志架构的大规模发布/订阅消息队列”，这使它作为企业级基础设施来处理流式数据非常有价值。

2. **Kafka 的性能**：Kafka 的写入性能非常高，因为它是为处理大量实时数据流设计的。Kafka 使用了一些优化技术，例如顺序写入和零拷贝，这使得它能够以非常高的吞吐量写入数据。此外，Kafka 的数据是分布式存储的，你可以通过增加更多的分区和副本来提高写入性能。

3. **使用 Protobuf 协议**：消息队列的读写使用 Protobuf 协议。Protocol Buffers（Protobuf）是 Google 开发的一种数据序列化协议（类似于 XML、JSON、YAML 等），它能够将结构化数据序列化，可用于数据存储、通信协议等方面。Protobuf 相比 JSON，XML 格式的数据，数据更小（3 到 10 倍的压缩率）、速度更快（20 到 100 倍的速度），并且 Protobuf 提供了丰富的数据结构，并且可以生成各种语言的数据访问代码，包括 Go。


## 主要功能
1. 待写[✅]  [❌] 

## pprof分析Go程序性能
```
go tool pprof http://localhost:6060/debug/pprof/heap  获取内存性能分析报告
go tool pprof http://localhost:6060/debug/pprof/profile 获取CPU性能分析报告
flat：函数直接分配的内存。  
flat%：函数直接分配的内存占总内存的百分比。  
sum%：到目前为止直接分配的内存占总内存的百分比。  
cum：函数及其所有子函数分配的内存。  
cum%：函数及其所有子函数分配的内存占总内存的百分比。

Commands:
    callgrind        输出callgrind格式的图表，这种格式可以被一些工具如KCachegrind或者QCacheGrind读取。
    comments         输出所有的分析注释。
    disasm           输出带有样本注释的汇编列表。
    dot              输出DOT格式的图表，这种格式可以被Graphviz等工具读取。
    eog              使用相应的工具可视化图表。eog
    evince           使用相应的工具可视化图表。 evince
    gif              输出相应格式的图表 GIF 
    gv               使用相应的工具可视化图表。 gv
    kcachegrind      在KCachegrind中可视化报告
    list             输出与正则表达式匹配的函数的带注释的源代码。
    pdf              输出相应格式的图表 PDF 
    peek             输出与正则表达式匹配的函数的调用者/被调用者。
    png              输出相应格式的图表 PNG 
    proto            输出压缩的protobuf格式的分析数据。
    ps               输出相应格式的图表 PS 
    raw              输出原始分析数据的文本表示。
    svg              输出相应格式的图表 SVG 
    tags             输出分析中的所有标签。
    text             以文本形式输出顶级条目。
    top              以文本形式输出顶级条目。
    topproto         以压缩的protobuf格式输出顶级条目。
    traces           以文本形式输出所有分析样本。
    tree             输出调用图的文本表示。
    web              在网页浏览器中可视化图表或显示带注释的源代码。
    weblist          在网页浏览器中可视化图表或显示带注释的源代码。
    o/options        列出所有选项及其当前值。
    q/quit/exit/^D   退出pprof
```




 