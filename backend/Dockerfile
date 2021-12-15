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

# 本番環境
FROM golang:1.15.2 as builder

WORKDIR /build
COPY . /build/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

FROM alpine:3.9.4

LABEL environment="production"
WORKDIR /app
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
EXPOSE 3000
ENTRYPOINT ["./main"]
