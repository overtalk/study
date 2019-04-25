# build go program first
FROM golang

WORKDIR /go/src/k8s/

ADD ./. .

RUN /usr/local/go/bin/go build  -o ./pvp ./server/pvp/server.go

ENV PVP="PVP_SERVER"

EXPOSE 8081

CMD ./pvp
