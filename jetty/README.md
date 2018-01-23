基于 https://github.com/appropriate/docker-jetty 修改：

- ENV JETTY_VERSION 9.4.8.v20171121

- --add-to-start="...,cdi2" --approve-all-licenses \

- RUN mkdir -p "$JETTY_BASE" 改为 RUN mkdir -p "$JETTY_BASE/lib/ext/jndi/"

- java -jar...前增加：
RUN set -xe \
	&& curl -SL "http://central.maven.org/maven2/org/mariadb/jdbc/mariadb-java-client/2.2.1/mariadb-java-client-2.2.1.jar" -o lib/ext/jndi/mariadb-java-client-2.2.1.jar


<br />

Usages：

mkdir -p ~/build-images/jetty/; cd ~/build-images/jetty/;

wget https://raw.githubusercontent.com/iwangxiaodong/container-config/master/jetty/Dockerfile

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/docker-entrypoint.sh

wget https://raw.githubusercontent.com/appropriate/docker-jetty/master/9.4-jre8/generate-jetty-start.sh

sudo chmod +x *

sudo docker build -t jetty-image .  # 若传docker仓库则 docker build -t localhost:55000/jetty .

sudo docker run -d -p 8088:8080 jetty-image  # curl http://localhost:8088

<br />

[若运行时修改jetty模块则要删除jetty.start并重启生效] sudo docker exec -it java-web-edu sh -c "rm jetty.start"

