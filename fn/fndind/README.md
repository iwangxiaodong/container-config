Usages：

mkdir -p ~/build-images/jetty/; cd ~/build-images/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/fndind/Dockerfile

sudo docker build -t my-fndind . # 若传docker仓库则 docker build -t localhost:55000/my-fndind .