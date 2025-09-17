#!/bin/bash

# 如果任何命令失败，脚本将立即退出
set -e

# 获取脚本所在的目录
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

echo ">>> 1. Building Vue.js frontend..."
cd "$SCRIPT_DIR/web"
if [ ! -d "node_modules" ]; then
  echo "Node modules not found, running npm install..."
  npm install
fi
npm run build
cd "$SCRIPT_DIR"

echo ">>> 2. Building Go backend for Windows..."
# 如果 go.mod 不存在，则初始化模块
if [ ! -f "go.mod" ]; then
  echo "Go module not found, running go mod init..."
  go mod init video-collector
  go mod tidy
fi

# 设置环境变量进行交叉编译，并构建应用
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o video-collector-windows.exe .

echo ">>> 3. Build finished successfully!"
echo "Run 'video-collector-windows.exe' on a Windows machine to start the server."