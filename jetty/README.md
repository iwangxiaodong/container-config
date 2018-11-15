基于 https://github.com/appropriate/docker-jetty 修改：

- ENV JETTY_VERSION 9.4.13.v20181111

- --add-to-start="...,cdi2" --approve-all-licenses \  # 后续移除jsp和jstl; 若用到https可添加http2模块

- RUN mkdir -p "$JETTY_BASE" 改为 RUN mkdir -p "$JETTY_BASE/lib/ext/jndi/"

- java -jar...前增加：
	&& curl -SL "http://central.maven.org/maven2/org/mariadb/jdbc/mariadb-java-client/2.3.0/mariadb-java-client-2.3.0.jar" -o lib/ext/jndi/mariadb-java-client-2.3.0.jar \
	
- 若为FROM openjdk:10-slim 还需增加：
RUN apt-get update && apt-get install -y --no-install-recommends curl gnupg libfontconfig1 && rm -rf /var/lib/apt/lists/*

<br />
注意：可选安装包 iputils-ping vim 仅为方便调试。
<br />

Usages：

mkdir -p ~/build-images/jetty/; cd ~/build-images/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/jetty/Dockerfile

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/docker-entrypoint.sh

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/generate-jetty-start.sh

sudo chmod +x *

sudo docker build -t my-jetty .  
# 若传docker仓库则 
  sudo docker build -t localhost:30500/cr/all/my-jetty .
  
  sudo docker push localhost:30500/cr/all/my-jetty

sudo docker run -d -p 8088:8080 my-jetty  # curl http://localhost:8088

<br />

[若运行时修改jetty模块则要删除jetty.start并重启生效] sudo docker exec -it java-web-edu sh -c "rm jetty.start"


k8s添加Java选项示例：
```yaml
        env:
        - name: JAVA_OPTIONS
          value: "-XX:+UnlockExperimentalVMOptions -XX:+Use..."
```

