FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY configs /app/configs

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone  && \
    go build -ldflags="-s -w" -o /app/packer cmd/packer-server/packerserver.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/packer /app/packer
COPY --from=builder /app/configs /app/configs


CMD ["./packer", "-f", "configs/packer-server.yaml"]