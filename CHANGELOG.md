# CHANGELOG

## 1.0.0

这是一个agent的示例

## 1.0.1

修改打包程序增加epoch

## 1.0.2

增加manifest.yaml配置信息描述文件

## 1.0.3

修正manifest.yaml中的id字段缺少type的问题

## 1.0.4

修改配置文件，增加basic前缀

## 1.0.5

1. 修改打包流程，增加manifest-lint过程
2. 修正manifest.yaml中basic.log.target字段默认值设置问题

## 1.0.6

1. 修改命令行交互方式
2. 去除go1.17版本支持

## 1.0.7

修正deb和rpm包安装时注册系统服务的问题

## 1.0.8

1. 修正windows下的系统服务启动问题
2. 去除msi安装包

## 1.0.9

windows下的exe安装包支持静默卸载

## 1.0.10

1. 修改启动、停止、重启、注册、反注册系统服务失败时的状态码
2. 修正windows系统下未启动服务无法重启的问题

## 1.0.11

1. manifest.yaml中增加required字段
2. 修正golint问题

## 1.0.12

修改manifest.yaml中的basic.id字段类型为uuid

## 1.1.0

支持arm架构

## 1.2.0

新增日志文件上报功能