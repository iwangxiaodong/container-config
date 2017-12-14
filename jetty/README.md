mkdir -p ~/docker-build/jetty/; cd ~/docker-build/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/jetty/Dockerfile

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/docker-entrypoint.sh

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/generate-jetty-start.sh

sudo docker build -t my-jetty .  # 若传docker仓库则 docker build -t localhost:55000/jetty .
