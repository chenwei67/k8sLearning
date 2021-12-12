FROM golang:1.16
COPY ./ /data/
WORKDIR /data/

RUN go mod download -x
ENTRYPOINT ["go", "run", "main.go"]
