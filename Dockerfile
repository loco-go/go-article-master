FROM golang:1.17-alpine AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o hello ./main.go

FROM gcr.io/distroless/static
WORKDIR /
COPY --from=builder /workspace/hello /app

EXPOSE 8080
CMD ["/app"]