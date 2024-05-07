# 一键安装脚本
version='v1.0.2'

# 下载
wget "https://github.com/mouday/shell-runner/releases/download/${version}/shell-runner-${version}-linux-amd64.tar.gz"

# 解压
tar -zxf "shell-runner-${version}-linux-amd64.tar.gz"

# 重命名
mv "shell-runner-${version}-linux-amd64" shell-runner

# 进入目录
cd shell-runner

# 创建.env文件
cp env.example .env

# 安装服务
cp ./config/shell-runner.service /etc/systemd/system/

sed -i "s@/opt/shell-runner@${$(pwd)}@g" /etc/systemd/system/shell-runner.service

# 开机自启
systemctl enable shell-runner

# 启动
systemctl start shell-runner
