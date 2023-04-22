FROM golang:1.18-alpine AS builder
ENV GOPROXY='https://goproxy.cn' \
    GOPRIVATE='xxx.xxx.com' \
    CGO_ENABLED=0
# ARG TOKEN
# RUN echo "machine xxx.xxx.com login anonymous password ${TOKEN}" > ~/.netrc
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk upgrade && apk add --no-cache git
WORKDIR /src
RUN mkdir ./bin
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./bin/main ./cmd/

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk upgrade
WORKDIR /app
RUN mkdir ./bin
COPY --from=builder /src/bin ./bin
COPY ./resource ./resource
# http port
EXPOSE 8080
# grpc port
EXPOSE 8081

ENTRYPOINT ["./bin/main"]
