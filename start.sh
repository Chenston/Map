#!/bin/bash

# 启动 MongoDB 容器
docker run --name mongo -d -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=admin123 mongo

# 等待 MongoDB 启动
echo "Waiting for MongoDB to start..."
sleep 5

# 导入数据到 MongoDB
docker cp states.json mongo:/states.json
docker exec -it mongo mongoimport --db us_states --collection states \
    --file /states.json --jsonArray --username admin \
    --password admin123 --authenticationDatabase admin

# 启动后端（新终端窗口）
if command -v gnome-terminal >/dev/null 2>&1; then
    gnome-terminal -- bash -c "cd backend && go run main.go; exec bash"
elif command -v konsole >/dev/null 2>&1; then
    konsole --noclose -e bash -c "cd backend && go run main.go"
elif command -v x-terminal-emulator >/dev/null 2>&1; then
    x-terminal-emulator -e bash -c "cd backend && go run main.go"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    open -a Terminal "backend" && osascript -e 'tell application "Terminal" to do script "cd backend && go run main.go"'
else
    echo "No compatible terminal found, running in the background..."
    (cd backend && go run main.go) &
fi

# 启动前端（新终端窗口）
if command -v gnome-terminal >/dev/null 2>&1; then
    gnome-terminal -- bash -c "cd frontend && npm install && npm run dev; exec bash"
elif command -v konsole >/dev/null 2>&1; then
    konsole --noclose -e bash -c "cd frontend && npm install && npm run dev"
elif command -v x-terminal-emulator >/dev/null 2>&1; then
    x-terminal-emulator -e bash -c "cd frontend && npm install && npm run dev"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    open -a Terminal "frontend" && osascript -e 'tell application "Terminal" to do script "cd frontend && npm install && npm run dev"'
else
    echo "No compatible terminal found, running in the background..."
    (cd frontend && npm install && npm run dev) &
fi

# 等待前端启动
sleep 3

# 自动打开浏览器访问前端页面
URL="http://localhost:5173/"
if command -v xdg-open >/dev/null 2>&1; then
    xdg-open "$URL"  # Linux
elif command -v open >/dev/null 2>&1; then
    open "$URL"  # macOS
else
    echo "Please open $URL manually"
fi

echo "All services started successfully!"
