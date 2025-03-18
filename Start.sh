#!/bin/bash

# 启动 MongoDB 容器
echo "Starting MongoDB container..."
docker run -d --name mongo -p 27017:27017 mongo

# 等待 MongoDB 启动
sleep 5
echo "MongoDB started."

# 填充数据库
echo "Importing states.json into MongoDB..."
docker exec -i mongo mongoimport --db us_states --collection states --file /states.json --jsonArray

# 进入后端目录并启动后端服务
echo "Starting backend service..."
cd backend
go run main.go &

# 等待后端启动
sleep 5
echo "Backend started."

# 进入前端目录
echo "Setting up frontend..."
cd ../frontend
npm install

# 启动前端
echo "Starting frontend service..."
npm run dev

# 等待前端启动
sleep 5
echo "Frontend started."

# 打开前端页面
echo "Opening frontend in browser..."
x-www-browser http://localhost:5173/ &

echo "All services started successfully."
