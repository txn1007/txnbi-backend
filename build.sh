#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# =======================
# 配置部分
# =======================

# 默认的可执行文件名称
BINARY_NAME="myapp"

# 输出目录，默认为当前目录
OUTPUT_DIR="$(pwd)"

# 要编译的包路径，默认为当前目录
PACKAGE_PATH="./cmd/webService"

# Go 环境变量（可选，用于跨平台编译）
# 例如：GOOS=linux GOARCH=amd64 ./build.sh
GOOS=${GOOS:-$(go env GOOS)}
GOARCH=${GOARCH:-$(go env GOARCH)}

# =======================
# 函数定义
# =======================

# 显示使用说明
usage() {
    echo "Usage: $0 [-n binary_name] [-o output_dir] [GOOS=linux GOARCH=amd64]"
    echo ""
    echo "Options:"
    echo "  -n    设置可执行文件的名称（默认：myapp）"
    echo "  -o    设置可执行文件的输出目录（默认：当前目录）"
    echo "  -h    显示帮助信息"
    exit 1
}

# 解析命令行参数
while getopts "n:o:h" opt; do
    case "$opt" in
        n) BINARY_NAME="$OPTARG" ;;
        o) OUTPUT_DIR="$OPTARG" ;;
        h|*) usage ;;
    esac
done

shift $((OPTIND-1))

# 检查 Go 是否已安装
if ! command -v go &> /dev/null
then
    echo "错误：Go 未安装。请访问 https://golang.org/dl/ 安装 Go。"
    exit 1
fi

# 检查 Go 版本（可选）
REQUIRED_GO_VERSION="1.23.2"
INSTALLED_GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')

if [[ "$INSTALLED_GO_VERSION" < "$REQUIRED_GO_VERSION" ]]; then
    echo "错误：Go 版本需要至少为 $REQUIRED_GO_VERSION。当前版本：$INSTALLED_GO_VERSION"
    exit 1
fi

# 创建输出目录（如果不存在）
mkdir -p "$OUTPUT_DIR"

# 编译项目
echo "开始编译项目..."
echo "GOOS: $GOOS"
echo "GOARCH: $GOARCH"
echo "输出目录: $OUTPUT_DIR"
echo "可执行文件名称: $BINARY_NAME"

go build -o "$OUTPUT_DIR/$BINARY_NAME" "$PACKAGE_PATH"

# 检查编译是否成功
if [ $? -eq 0 ]; then
    echo "编译成功！可执行文件位于：$OUTPUT_DIR/$BINARY_NAME"
else
    echo "编译失败！"
    exit 1
fi
