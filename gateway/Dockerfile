FROM golang:1.19
WORKDIR /mint-services/src/gateway
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]