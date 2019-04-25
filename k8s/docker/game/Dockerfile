# build go program first
FROM golang

WORKDIR /go/src/k8s/

ADD ./. .

RUN /usr/local/go/bin/go build  -o ./game ./server/game/server.go

ENV GAME="game_server"

EXPOSE 8080

CMD ./game
