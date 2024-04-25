# Shell Runner 

Shell Runner 是一个基于Go语言编写的shell脚本执行工具，通过http请求触发shell脚本异步执行。

使用场景

1. 可配合代码仓库的代码推送事件，通过webhook触发代码打包编译使用
2. 通过http请求远程执行某些命令，例如：定时任务

## 安装

1、下载适合所在平台的可执行文件

[https://github.com/mouday/shell-runner/releases](https://github.com/mouday/shell-runner/releases)


2、配置环境变量

```bash
cp env.example .env
```

配置文件`.env`

```bash
# == 应用配置 ==
# 运行模式 debug test release (默认：release)
GIN_MODE=release

# 监听端口 (默认：127.0.0.1:8000）
APP_RUN_ADDRESS=127.0.0.1:8000

# 权限 token 必填，若为空，则无法调用接口
APP_TOKEN=

# 脚本存放目录
APP_SCRIPT_DIR=./scripts
```

3、启动

```bash
# macos: 
./shell-runner

# linux: 
./shell-runner

# windows: 
shell-runner.exe
```

## 使用


例如：

在目录`scripts` 下有一个脚本`hello.sh`

```bash
./scripts
    /hello.sh
```

可以通过http请求执行

```bash
POST http://localhost:8000/run?name=hello
Content-Type: application/json; charset=utf-8
X-Token: <token>

```

参数说明

| 参数名 | 是否必填 | 说明 |
| --- | --- | --- |
| name | 是 | 脚本名称，不带后缀 |
| token | 是 | 权限token，若为空，则无法调用接口 |

> 注意：token必须填写，和环境变量中配置一致


## 通过systemd让进程自动启动

略
