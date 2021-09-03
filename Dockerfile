#源镜像
FROM ubuntu
#作者
MAINTAINER ZhangLianjun "2814634354@qq.com"

# 修改国内源
RUN sed -i s/archive.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list && sed -i s/security.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list
RUN apt-get update && apt-get upgrade -y
RUN apt-get install gcc libc6-dev git lrzsz inetutils-ping curl -y
#禁止交互
ENV DEBIAN_FRONTEND=noninteractive

# 安装Go环境
RUN apt-get install golang -y

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on


#工作目录作目录
WORKDIR $GOPATH/src/unioj

COPY . .
ADD go.mod .
ADD go.sum .

# 自己构建docker的是否注意删掉这三句  以便镜像能找到beego
RUN cp /etc/hosts ~/hosts.new
RUN sed -i '2i\192.168.10.63   unioj.org' ~/hosts.new
RUN cp -f ~/hosts.new /etc/hosts

RUN rm -rf ~/hosts.new

RUN go mod download

#拷贝文件
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
WORKDIR $GOPATH/src/unioj/run
RUN go build run.go
#暴露端口
EXPOSE 7999
#最终运行docker的命令
CMD  ["./run"]

#   docker build -t unioj-judger .
#   docker run --name unioj-judger -p 7999:7999 --link unioj-beego  `docker images | grep  'unioj-judger' | awk '{print $3}'`
#   docker run --name unioj-judger -p 7998:7999   b93ac098bfc9
#一条命令实现停用并删除容器：
#   docker stop $(docker ps -q) & docker rm $(docker ps -aq)
#删除none镜像
#   docker rmi `docker images | grep  '<none>' | awk '{print $3}'`

#   docker tag 5f51660b5491 registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-judger:latest
#   docker push registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-judger:latest