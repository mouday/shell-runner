# Shell Runner 

使用方法

安装之后通过http请求触发异步执行

例如：

在目录`scripts` 下有一个脚本`hello.sh`

```bash
/scripts
    /hello.sh
```

可以通过http请求执行

```bash
POST http://localhost:8000/run?name=hello
Content-Type: application/json; charset=utf-8
X-Token: <token>

```

> 注意：token必须填写，在环境变量中配置

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