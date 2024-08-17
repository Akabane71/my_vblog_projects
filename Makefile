# 这些命令都是
.PHONY: all build run gotool clean help
# 定义一下编译后的文件名
BINARY="bluebell"
#  定义目标
all: gotool build
# 编译xx目标
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go ./conf/config.yaml

gotool:
	go fmt ./
	go vet ./

# 支持写shell命令的
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
