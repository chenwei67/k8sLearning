FROM golang:1.16-buster
COPY aaa /data/
WORKDIR /data/

RUN go mod download -x
ENTRYPOINT ["go", "run", "main.go"]
