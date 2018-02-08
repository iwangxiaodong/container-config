Usages：

mkdir -p ~/build-images/fndind/; cd ~/build-images/fndind/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/fn/fndind/Dockerfile

sudo docker build -t my-fndind . # 若传docker仓库则 docker build -t localhost:30500/my-fndind .
