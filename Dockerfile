FROM golang:1.20.2-buster AS dev

RUN  wget -q -nv -O /bin/hadolint https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
    chmod 777 /bin/hadolint

WORKDIR /home/app
COPY . .

ENV HTTP_PORT=8080 \
    LOG_LEVEL=debug

RUN go install github.com/cosmtrek/air@v1.42.0 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2 && \
    go mod download

CMD ["air", "-c", ".air.toml"]