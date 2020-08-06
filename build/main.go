FROM alpine:latest
RUN apk update && \
  apk add git && \
  apk add go && \ 
  wget -O /root/caddy.tar.gz \
    "https://github.com/GooseG17/caddy/archive/v1.zip" && \
  tar -xf /root/caddy.tar.gz && \
  mv /root/caddy-1/ /root/caddy/ && \
  mkdir /root/caddy/build && \
  touch /root/caddy/build/build.go && \
  
