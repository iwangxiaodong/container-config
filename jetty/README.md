基于 https://github.com/appropriate/docker-jetty 修改：

- ENV JETTY_VERSION 9.4.8.v20171121

- --add-to-start="...,cdi2" --approve-all-licenses \

<br />

Usages：

mkdir -p ~/docker-build/jetty/; cd ~/docker-build/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/jetty/Dockerfile

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/docker-entrypoint.sh

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/generate-jetty-start.sh

sudo chmod +x *

sudo docker build -t my-jetty .  # 若传docker仓库则 docker build -t localhost:55000/jetty .

sudo docker run -d -p 8088:8080 my-jetty  # curl http://localhost:8088

<br />

[若运行时修改jetty模块则要删除jetty.start并重启生效] sudo docker exec -it java-web-edu sh -c "rm jetty.start"

