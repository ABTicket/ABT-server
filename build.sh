#!/bin/bash

# 从github上拉取最新的代码
git pull origin dev

# 停止容器
sudo docker stop ticket-server
# 更新代码至容器中
sudo docker cp . ticket-server:/app
# 启动容器
sudo docker start ticket-server

echo "Starting service..."
# 在容器中启动服务
sudo docker exec ticket-server /bin/bash -c \
    "cd src && go build main.go && ./main >out.log 2>&1 &"
echo "Deploy done!"
