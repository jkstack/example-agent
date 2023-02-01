# example-agent

[![version](https://img.shields.io/github/v/release/jkstack/example-agent)](https://github.com/jkstack/example-agent/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkstack/example-agent)](https://goreportcard.com/report/github.com/jkstack/example-agent)
[![go-mod](https://img.shields.io/github/go-mod/go-version/jkstack/example-agent)](https://github.com/jkstack/example-agent)
[![license](https://img.shields.io/github/license/jkstack/example-agent)](https://opensource.org/licenses/MIT)
![downloads](https://img.shields.io/github/downloads/jkstack/example-agent/total)

这是一个agent的示例，用于展示agent开发

## agent开发

1. 创建项目并导入[libagent](github.com/jkstack/libagent)库

    ```
    go get github.com/jkstack/libagent
    ```
2. 定义一个结构体并实现[App](https://pkg.go.dev/github.com/jkstack/libagent#App)接口
3. 在main函数中添加RegisterService、UnregisterService和Run等libagent库中的接口调用

    ```go
    switch *act {
    case "install":
        agent.RegisterService(app)
    case "uninstall":
        agent.UnregisterService(app)
    default:
        agent.Run(app)
    }
    ```

## 更多示例

1. [metrics-agent](https://github.com/jkstack/metrics-agent): 监控数据采集
2. [exec-agent](https://github.com/jkstack/exec-agent): 远程执行脚本，文件分发

## 开源社区

文档知识库：https://jkstack.github.io/

<img src="docs/wechat_QR.jpg" height=200px width=200px />