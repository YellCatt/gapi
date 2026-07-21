#!/bin/bash
set -e

echo "=== 开始构建 MIPSle 版本 ==="

# 检查交叉编译工具链
if ! command -v mipsel-linux-gnu-gcc &> /dev/null; then
    echo "错误: 未找到 mipsel-linux-gnu-gcc，请先安装交叉编译工具链"
    echo "Ubuntu/Debian: sudo apt-get install gcc-mipsel-linux-gnu"
    exit 1
fi

# 版本信息
SQLITE_VERSION="3450300"
BUILD_DIR="build_mipsle"
LIB_DIR="$BUILD_DIR/lib"
OUTPUT="gapi_linux_mipsle"

# 创建目录
mkdir -p "$BUILD_DIR"
mkdir -p "$LIB_DIR"

# 下载并编译 SQLite
echo "1. 下载 SQLite 源码..."
if [ ! -f "sqlite-autoconf-${SQLITE_VERSION}.tar.gz" ]; then
    wget -q https://www.sqlite.org/2024/sqlite-autoconf-${SQLITE_VERSION}.tar.gz
fi

echo "2. 解压 SQLite..."
tar -xzf sqlite-autoconf-${SQLITE_VERSION}.tar.gz

echo "3. 交叉编译 SQLite..."
cd sqlite-autoconf-${SQLITE_VERSION}
./configure --host=mipsel-linux-gnu --prefix=$(pwd)/../${LIB_DIR} --enable-shared
make -j$(nproc)
make install
cd ..

echo "4. 清理 SQLite 源码..."
rm -rf sqlite-autoconf-${SQLITE_VERSION}

# 设置环境变量
export CC=mipsel-linux-gnu-gcc
export CXX=mipsel-linux-gnu-g++
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=mipsle
export GOMIPS=softfloat

echo "5. 编译 Go 程序..."
go build \
    -ldflags="-s -w -linkmode=external -extldflags='-Wl,-rpath=\$ORIGIN/lib'" \
    -tags="netgo osusergo" \
    -o "$BUILD_DIR/$OUTPUT"

echo "6. 复制动态库..."
cp ${LIB_DIR}/lib/*.so* ${BUILD_DIR}/lib/

echo ""
echo "=== 构建完成 ==="
echo "输出目录: $(pwd)/${BUILD_DIR}"
echo "运行方式: ./${OUTPUT}"
echo ""
echo "部署结构:"
echo "  ${BUILD_DIR}/"
echo "    ├── ${OUTPUT}      # 可执行文件"
echo "    └── lib/           # SQLite 动态库"
echo "        └── libsqlite3.so.0"