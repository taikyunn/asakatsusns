# 本番環境
FROM golang:1.15.2 as builder

WORKDIR /build
COPY ./backend /build/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

FROM alpine:3.9.4

LABEL environment="production"
WORKDIR /app
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
EXPOSE 3000
ENTRYPOINT ["./main"]
