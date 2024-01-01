# 現在の HEAD にタグを付ける
VERSION := $(shell git describe --tags --abbrev=0)

# タグを含むコミットの SHA-1 ハッシュ値を取得
COMMIT_HASH := $(shell git rev-parse --short HEAD)

# バージョン名にコミットハッシュ値を付加
VERSION := $(VERSION)-$(COMMIT_HASH)

OPENCV_RUNTIME := C:\\opencv\\build\\install\\x64\\mingw\\bin

build:
	go build .

run: build
	./wrc-logger.exe

pack: build sync
	pack.bat wrc-logger-$(VERSION).zip
	#powershell Compress-Archive -Path dist -Force -DestinationPath wrc-logger-$(VERSION).zip

version:
	@echo $(VERSION) | cat

sync:
	robocopy dist\log log /E /XO