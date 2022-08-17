# example-agent

[![example-agent](https://github.com/jkstack/example-agent/actions/workflows/build.yml/badge.svg)](https://github.com/jkstack/example-agent/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkstack/example-agent)](https://goreportcard.com/report/github.com/jkstack/example-agent)
[![go-mod](https://img.shields.io/github/go-mod/go-version/jkstack/example-agent)](https://github.com/jkstack/example-agent)
[![license](https://img.shields.io/github/license/jkstack/example-agent)](https://opensource.org/licenses/MIT)

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

TODO