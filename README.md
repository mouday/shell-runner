# Shell Runner 

Shell Runner 是一个基于Go语言编写的shell脚本执行工具，通过http请求触发shell脚本异步执行。

使用场景

1. 可配合代码仓库的代码推送事件，通过webhook触发代码打包编译使用
2. 通过http请求远程执行某些命令，例如：定时任务

目的是替换jenkins的一部分功能，jenkins需要在java的环境下才能运行

依赖 | 大小
--- | ---
OpenJDK8U-jdk_aarch64_linux_hotspot_8u402b06.tar.gz | 97.4 MB
jenkins-war-stable-2.346.3.war | 86.97 MB


可以看到，整个运行环境大小接近200MB，而Shell Runner 不到10MB

## 安装

1、下载适合所在平台的可执行文件

[https://github.com/mouday/shell-runner/releases](https://github.com/mouday/shell-runner/releases)

以`v1.0.2`版本为例

```bash
# 下载
wget https://github.com/mouday/shell-runner/releases/download/v1.0.2/shell-runner-v1.0.2-linux-amd64.tar.gz

# 解压
tar -zxvf shell-runner-v1.0.2-linux-amd64.tar.gz

# 重命名
mv shell-runner-v1.0.2-linux-amd64 shell-runner

# 进入目录
cd shell-runner-v1.0.2-linux-amd64
```


2、启动

```bash
# macos: 
./shell-runner

# linux: 
./shell-runner

# windows: 
shell-runner.exe
```

启动后会生成一个token，用作权限校验

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

查看token

```bash
# 查看token
cat token.txt
a8c2be92-9c53-42d1-ad67-ee67ea58a3ac
```

## 配置环境变量

```bash
cp env.example .env
```

配置文件`.env`

```bash
# == 应用配置 ==
# 运行模式，可选: debug/test/release (默认：release)
GIN_MODE=release

# 监听端口 (默认：127.0.0.1:8000）
APP_RUN_ADDRESS=127.0.0.1:8000

# 脚本存放目录 (默认：./scripts)
APP_SCRIPT_DIR=./scripts
```

## 通过systemd让进程自动启动

在linux系统下，可以通过systemd让进程自动启动

```bash

cp ./config/shell-runner.service /etc/systemd/system/

# 开机自启
systemctl enable shell-runner

# 启动
systemctl start shell-runner

# 查看状态
systemctl status shell-runner

# 查看日志
journalctl -u shell-runner -f
```

通常使用nginx来统一转发请求

nginx config

```bash
# nginx config
server
{
    listen 8001;

    server_name localhost;

    # log
    if ($time_iso8601 ~ '(\d{4}-\d{2}-\d{2})') {
        set $time $1;
    }

    access_log /data/wwwlogs/nginx_log/shell-runner_${time}.log main;

    # 代理服务器
    location / {
        proxy_pass http://127.0.0.1:8000/;
        include proxy.conf;
    }
}
```

proxy.conf 文件内容

```bash
proxy_connect_timeout 300s;
proxy_send_timeout 900;
proxy_read_timeout 900;
proxy_buffer_size 32k;
proxy_buffers 4 64k;
proxy_busy_buffers_size 128k;
proxy_redirect off;
proxy_hide_header Vary;
proxy_set_header Accept-Encoding '';
proxy_set_header Referer $http_referer;
proxy_set_header Cookie $http_cookie;
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
```