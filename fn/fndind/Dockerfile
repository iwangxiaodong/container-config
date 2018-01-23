FROM fnproject/dind
RUN apk add --no-cache ca-certificates wget
RUN wget https://github.com/fnproject/cli/releases/download/0.4.41/fn_alpine -O /usr/local/bin/fn
RUN chmod +x /usr/local/bin/fn
WORKDIR /app/data
ENTRYPOINT ["fn"]
