#源镜像
FROM openjdk:8-jdk-alpine
#作者
MAINTAINER ZhangLianjun "2814634354@qq.com"

#禁止交互
ENV DEBIAN_FRONTEND=noninteractive


#工作目录作目录
WORKDIR app

ADD oj-0.0.1-SNAPSHOT.jar  /oj-0.0.1-SNAPSHOT.jar
COPY . .
EXPOSE 8081
CMD java -jar /oj-0.0.1-SNAPSHOT.jar

#   docker build -t unioj-fileserver .
#   docker run --name unioj-fileserver -p 8081:8081   `docker images | grep  'unioj-fileserver' | awk '{print $3}'`
#   docker run --name unioj-fileserver -p 8081:8081   65863d7f91bd
#一条命令实现停用并删除容器：
#   docker stop $(docker ps -q) & docker rm $(docker ps -aq)
#删除none镜像
#   docker rmi `docker images | grep  '<none>' | awk '{print $3}'`

#   docker tag 5f51660b5491 registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-fileserver:latest
#   docker push registry.cn-beijing.aliyuncs.com/uyistcoj/unioj-fileserver:latest