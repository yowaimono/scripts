#!/bin/bash

# MySQL 容器的名称
CONTAINER_NAME="mysql_container"

# MySQL 的 root 密码
MYSQL_ROOT_PASSWORD="123ttk"

# 数据挂载目录
DATA_DIR="/home/ttk/zyh_work/data/mysql"

# 创建数据挂载目录
mkdir -p $DATA_DIR

# 运行 MySQL 容器
docker run -it -d -p 3306:3306 --name $CONTAINER_NAME \
    -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
    -v $DATA_DIR:/var/lib/mysql \
    -d m.daocloud.io/docker.io/library/mysql

echo "MySQL container started with name: $CONTAINER_NAME"
echo "Root password: $MYSQL_ROOT_PASSWORD"
echo "Data directory: $DATA_DIR"
