#源镜像
FROM golang:1.13
#作者
MAINTAINER ZhangLianjun "2814634354@qq.com"

#禁止交互
ENV DEBIAN_FRONTEND=noninteractive


ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on


#工作目录作目录
WORKDIR $GOPATH/src/unioj

COPY . .
ADD go.mod .
ADD go.sum .
RUN go mod download

#go构建可执行文件
RUN go build main.go
#暴露端口
EXPOSE 8000
#最终运行docker的命令
CMD  ["./main"]

#   docker build -t unioj-beego .
#   docker run --name unioj-beego -p 8000:8000 -p 8666:8666  `docker images | grep  'unioj-beego' | awk '{print $3}'`
#   一条命令实现停用并删除容器：
#   docker stop $(docker ps -q) & docker rm $(docker ps -aq)
#   删除none镜像
#docker rmi `docker images | grep  '<none>' | awk '{print $3}'`