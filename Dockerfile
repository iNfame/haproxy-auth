FROM haproxytech/haproxy-debian:latest

COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
COPY auth.lua /usr/local/etc/haproxy/auth.lua
COPY backend.map /usr/local/etc/haproxy/backend.map

EXPOSE 20100