Usages：

mkdir -p ~/build-images/fndind/; cd ~/build-images/fndind/;

wget https://raw.githubusercontent.com/docker-library/docker/master/18.02/dind/Dockerfile

wget https://raw.githubusercontent.com/docker-library/docker/master/18.02/dind/dockerd-entrypoint.sh

chmod 777 dockerd-entrypoint.sh

vim Dockerfile  # 添加到COPY ...前

- RUN wget https://github.com/fnproject/cli/releases/download/0.4.45/fn_alpine -O /usr/local/bin/fn

- RUN chmod +x /usr/local/bin/fn

- WORKDIR /app/data

sudo docker build -t my-fndind . 

# 若传docker仓库则 
  sudo docker build -t localhost:30500/cr/all/my-fndind .
  
  sudo docker push localhost:30500/cr/all/my-fndind


# 注意 - 需要-d参数
sudo docker run --privileged --name my-fndind -d localhost:30500/cr/all/my-fndind

sudo docker exec -it my-fndind sh -c 'docker version'

