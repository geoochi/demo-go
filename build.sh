#!/bin/bash

# 创建构建目录
mkdir -p dist

# 设置版本号（可以从git tag或环境变量获取）
VERSION="1.0.0"

# 构建函数
build() {
    os=$1
    arch=$2
    output="dist/server-${os}-${arch}-${VERSION}"

    # Windows需要加.exe后缀
    if [ "$os" = "windows" ]; then
        output="${output}.exe"
    fi

    echo "Building for $os/$arch..."
    GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o "$output" cmd/server/main.go

    if [ $? -eq 0 ]; then
        echo "✅ Successfully built for $os/$arch"
    else
        echo "❌ Failed to build for $os/$arch"
    fi
}

# 清理旧的构建文件
rm -rf dist/*

# 构建所有平台
# build linux amd64
# build linux arm64
# build darwin amd64
build darwin arm64
# build windows amd64

# 为每个二进制文件创建SHA256校验和
cd dist
echo -e "\nGenerating checksums..."
shasum -a 256 * >checksums.txt
cat checksums.txt
