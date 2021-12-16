# 開発環境
FROM golang:1.15.2-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go get -u github.com/cosmtrek/air
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 3000
CMD ["air", "-c", ".air.toml"]
