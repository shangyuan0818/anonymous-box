FROM golang:1.20-alpine AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

ARG SERVICE_NAME
ARG SERVICE_PATH=services/${SERVICE_NAME}/cmd/*.go

COPY . .

RUN go build -o /go/bin/service-entry ${SERVICE_PATH}

FROM alpine:latest
WORKDIR /data

COPY --from=builder /go/bin/service-entry /go/bin/service-entry

ENTRYPOINT ["/go/bin/service-entry"]
