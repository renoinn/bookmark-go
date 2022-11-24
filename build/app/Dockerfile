FROM golang:1.19.3-alpine

ENV APP_ROOT /var/app

WORKDIR ${APP_ROOT}
COPY ./ ${APP_ROOT}

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
