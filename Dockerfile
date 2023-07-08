FROM golang:1.20.5-alpine3.17 as goalpine
LABEL maintainer="lulianom"

RUN apk update && apk add --no-cache git \
  openssh openssh-keygen gcc \
  openssl-libs-static  \
  zlib-static zstd-libs libsasl \
  lz4-dev lz4-static libc-dev musl-dev

FROM goalpine AS builder

ARG tag

WORKDIR /app
RUN echo "$tag" > /app/version.file
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOGOARCH=amd64 GOOS=linux go build -tags musl -ldflags '-w -extldflags "-static"' -a -installsuffix cgo -o main main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && apk -U upgrade
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/version.file .
EXPOSE 8080
CMD ["./main"] 