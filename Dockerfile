FROM golang:1.20.2-buster AS dev

WORKDIR /home/app
COPY . .

ENV HTTP_PORT=8080 \
    LOG_LEVEL=debug

RUN go install github.com/cosmtrek/air@v1.42.0 && \
    go mod download



CMD air -c .air.toml