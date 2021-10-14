#!/bin/sh

# 这个脚本用于一键部署整个oj项目
#安装docker
curl -sSL https://get.daocloud.io/docker | sh

#安装docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version

#安装unioj系统
wget -P / https://gitee.com/ethanalban/uyistc-oj/attach_files/850474/download/docker-compose.yml
cd /
docker-compose up