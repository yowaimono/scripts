#!/bin/bash

# 定义变量
REDIS_IMAGE="m.daocloud.io/docker.io/library/redis"
REDIS_PORT=6379
DATA_DIR="/home/ttk/zyh_work/data/cache"
CONTAINER_NAME="redis_container"

# 检查数据目录是否存在，如果不存在则创建
if [ ! -d "$DATA_DIR" ]; then
  mkdir -p "$DATA_DIR"
  echo "Created data directory: $DATA_DIR"
fi

# 启动 Redis 容器
docker run -d \
  --name $CONTAINER_NAME \
  -p $REDIS_PORT:6379 \
  -v $DATA_DIR:/data \
  $REDIS_IMAGE

# 检查容器是否启动成功
if [ $? -eq 0 ]; then
  echo "Redis container started successfully!"
  echo "Access Redis on localhost:$REDIS_PORT"
else
  echo "Failed to start Redis container."
fi
