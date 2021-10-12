#源镜像
FROM ubuntu
#作者
MAINTAINER ZhangLianjun "2814634354@qq.com"
#禁止交互
ENV DEBIAN_FRONTEND=noninteractive
# 修改国内源
#ADD sources.list /etc/apt
RUN sed -i s/archive.ubuntu.com/mirrors.aliyun.com/ /etc/apt/sources.list && sed -i s/security.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list
RUN apt-get update && apt-get upgrade -y
RUN apt-get install gcc -y
RUN apt-get install build-essential -y
RUN apt-get install libc6-dev -y
RUN apt-get install git -y
RUN apt-get install lrzsz -y
RUN apt-get install inetutils-ping -y
RUN apt-get install curl -y
RUN apt-get install openjdk-8-jdk -y
RUN apt-get install netcat -y
#禁止交互
ENV DEBIAN_FRONTEND=noninteractive

# 安装Go环境
RUN apt-get install golang -y

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV JUDGERIP 47.96.99.172:7999

#工作目录作目录
WORKDIR $GOPATH/src/unioj-judger

COPY . .
ADD go.mod .
ADD go.sum .

# 自己构建docker的是否注意删掉这三句  以便镜像能找到beego
RUN cp /etc/hosts ~/hosts.new
RUN sed -i '2i\192.168.10.63   unioj.org' ~/hosts.new
RUN cp -f ~/hosts.new /etc/hosts
RUN chmod +x $GOPATH/src/unioj-judger/server_detect.sh
RUN chmod +x $GOPATH/src/unioj-judger/run/run


RUN rm -rf ~/hosts.new

RUN go mod download

#拷贝文件
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
WORKDIR $GOPATH/src/unioj-judger/run
RUN go build run.go
#暴露端口
EXPOSE 7999
EXPOSE 8081
#最终运行docker的命令
CMD  ["./run"]

#   docker build -t unioj-judger .
#   docker run --name unioj-judger -p 7999:7999 -p 8081:8081  --link unioj-beego  `docker images | grep  'unioj-judger' | awk '{print $3}'`
#   docker run --name unioj-judger -p 7999:7999 -p 8081:8081  --link unioj-beego 4a1553180c59
#   docker run --name unioj-judger -p 7998:7999   b93ac098bfc9
#一条命令实现停用并删除容器：
#   docker stop $(docker ps -q) & docker rm $(docker ps -aq)
#删除none镜像
#   docker rmi `docker images | grep  '<none>' | awk '{print $3}'`

#   docker tag 73c2bce5bdce registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-judger:latest
#   docker push registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-judger:latest