# 本番環境
FROM golang:1.15.2

WORKDIR /build
COPY ./backend /build/
RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

EXPOSE 3000
ENTRYPOINT ["./main"]
