FROM golang:alpine AS builder
ARG VERSION
RUN apk add --no-cache --virtual .build-deps git libc6-compat build-base
WORKDIR /system-monitor

COPY . .
RUN go mod download
WORKDIR /system-monitor/cmd/svr
RUN go build -gcflags "all=-N -l" -ldflags "-X main.version=${VERSION}" -o app;

FROM alpine:latest
WORKDIR /root
COPY --from=builder /system-monitor/cmd/svr/app .

CMD ["./app"]
