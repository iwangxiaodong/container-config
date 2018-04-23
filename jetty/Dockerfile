FROM openjdk:10-slim

# add our user and group first to make sure their IDs get assigned consistently, regardless of whatever dependencies get added
RUN groupadd -r jetty && useradd -r -g jetty jetty

ENV JETTY_HOME /usr/local/jetty
ENV PATH $JETTY_HOME/bin:$PATH
RUN mkdir -p "$JETTY_HOME"
WORKDIR $JETTY_HOME

ENV JETTY_VERSION 9.4.9.v20180320
ENV JETTY_TGZ_URL https://repo1.maven.org/maven2/org/eclipse/jetty/jetty-home/$JETTY_VERSION/jetty-home-$JETTY_VERSION.tar.gz
	
RUN apt-get update && apt-get install -y --no-install-recommends \
		curl libfontconfig1 \
	&& rm -rf /var/lib/apt/lists/*

RUN set -xe \
	&& curl -SL "$JETTY_TGZ_URL" -o jetty.tar.gz \
	&& tar -xvf jetty.tar.gz --strip-components=1 \
	&& sed -i '/jetty-logging/d' etc/jetty.conf \
	&& rm jetty.tar.gz* \
	&& rm -rf /tmp/hsperfdata_root

ENV JETTY_BASE /var/lib/jetty
RUN mkdir -p "$JETTY_BASE/lib/ext/jndi/"
WORKDIR $JETTY_BASE

RUN set -xe \
	&& curl -SL "http://central.maven.org/maven2/org/mariadb/jdbc/mariadb-java-client/2.2.3/mariadb-java-client-2.2.3.jar" -o lib/ext/jndi/mariadb-java-client.jar \
	&& curl -SL "http://central.maven.org/maven2/com/sun/activation/javax.activation/1.2.0/javax.activation-1.2.0.jar" -o lib/ext/javax.activation.jar \
	&& java -jar "$JETTY_HOME/start.jar" --create-startd --add-to-start="server,http,deploy,jsp,jstl,ext,resources,websocket,cdi2" --approve-all-licenses \
	&& chown -R jetty:jetty "$JETTY_BASE" \
	&& rm -rf /tmp/hsperfdata_root

ENV TMPDIR /tmp/jetty
RUN set -xe \
	&& mkdir -p "$TMPDIR" \
	&& chown -R jetty:jetty "$TMPDIR"

COPY docker-entrypoint.sh generate-jetty-start.sh /

USER jetty
EXPOSE 8080
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["java","-jar","/usr/local/jetty/start.jar"]
