FROM golang:1.20-alpine AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/migrator database/migrator/*.go

FROM alpine:latest
WORKDIR /data

COPY --from=builder /go/bin/migrator /go/bin/migrator

ENTRYPOINT ["/go/bin/migrator"]
