FROM golang:1.19

WORKDIR /mint/mint-services

COPY .. .

RUN go build -o bin/server gateway/cmd/main.go

CMD ["./bin/server"]