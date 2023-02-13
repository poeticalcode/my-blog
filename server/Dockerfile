FROM golang:1.19.3  

# 配置环境变量
ENV GOPROXY https://goproxy.cn,direct \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .
# 将我们的代码编译成二进制可执行文件 app
RUN go build -o main ./main.go
# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist
# 将二进制文件从 /build 目录复制到这里
RUN cp /build/main .

EXPOSE 9090

CMD [ "/dist/main" ]
