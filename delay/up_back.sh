#!/bin/bash

# 配置参数
BACKEND_DIR="../cmd/"    # 后端构建路径
SERVER_USER="***"        # 服务器用户名
SERVER_IP="***.**.***.**" # 服务器 IP 地址
SERVER_BACKEND_PATH="/home/ttk/zyh_work/mjadmin" # 服务器前端路径

# 打包前端项目
cd $BACKEND_DIR

# 配置交叉编译环境为 Linux
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0  # 禁用 cgo

# 编译项目
go build -o main main.go

# 上传后端项目
scp -r main $SERVER_USER@$SERVER_IP:$SERVER_BACKEND_PATH

# 清理环境变量
unset GOOS
unset GOARCH
unset CGO_ENABLED

echo "上传完成！"