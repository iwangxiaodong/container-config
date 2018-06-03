FROM ubuntu:latest
LABEL maintainer="Piasy Xu (xz4215@gmail.com)"

EXPOSE 8080 8089 3478 3033 59000-65000

WORKDIR /

ENV GAE_VER 1.9.70
#ENV GOLANG_VER 1.8.3

ENV PUBLIC_IP 127.0.0.1

RUN apt-get update -y

# Deps
RUN apt-get install -y build-essential vim git wget unzip python2.7 python-pil python-webtest python-pip libssl-dev
RUN pip install requests

# AppRTC
WORKDIR /
RUN git clone https://github.com/webrtc/apprtc.git

WORKDIR apprtc

RUN apt-get install -y golang
ENV GOPATH /goWorkspace
RUN mkdir -p $GOPATH/src
RUN go env

RUN ln -s /apprtc/src/collider/collider $GOPATH/src
RUN ln -s /apprtc/src/collider/collidermain $GOPATH/src
RUN ln -s /apprtc/src/collider/collidertest $GOPATH/src
RUN go get collidermain
RUN go install collidermain

RUN apt-get install -y nodejs
RUN apt install -y npm openjdk-8-jre
#RUN npm install -g npm
RUN npm install -g grunt-cli

RUN npm install
RUN grunt build

WORKDIR /

# Google App Engine
ENV GAE_ZIP google_appengine_$GAE_VER.zip
RUN wget https://storage.googleapis.com/appengine-sdks/featured/$GAE_ZIP
RUN unzip $GAE_ZIP -d /usr/local
ENV PATH $PATH:/usr/local/google_appengine

# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY run.sh /
RUN chmod +x /run.sh
CMD /run.sh
