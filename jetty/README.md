mkdir -p ~/docker-build/jetty/; cd ~/docker-build/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/jetty/Dockerfile

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/docker-entrypoint.sh

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/generate-jetty-start.sh

sudo chmod +x *

sudo docker build -t my-jetty .  # 若传docker仓库则 docker build -t localhost:55000/jetty .
