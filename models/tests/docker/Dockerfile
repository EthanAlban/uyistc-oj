#源镜像
FROM ubuntu:16.04
#作者
MAINTAINER ZhangLianjun "2814634354@qq.com"

#禁止交互
ENV DEBIAN_FRONTEND=noninteractive

RUN sed -i s/archive.ubuntu.com/mirrors.aliyun.com/ /etc/apt/sources.list && sed -i s/security.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list
RUN apt-get update && apt-get upgrade -y
RUN apt-get install inetutils-ping -y
# 安装Go环境
RUN apt-get install golang -y

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV name zhang_lianjun

#工作目录作目录
WORKDIR $GOPATH/src/unioj

COPY . .

#go构建可执行文件
RUN go build test.go
#暴露端口
#EXPOSE 8666
#EXPOSE 8000
#最终运行docker的命令
CMD  ["./test"]

#   docker build -t  edge-test .
#   docker run --name  edge-test   e13df055de9d
#   一条命令实现停用并删除容器：
#   docker stop $(docker ps -q) & docker rm $(docker ps -aq)
#   删除none镜像
#   docker rmi `docker images | grep  '<none>' | awk '{print $3}'`

#   docker tag c2d0e07b8338 registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-beego:latest
#   docker push registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-beego:latest
