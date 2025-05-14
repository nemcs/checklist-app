FROM golang:1.24.2-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext

COPY ["api-service/go.mod", "api-service/go.sum", "./api-service/"]
COPY ["db-service/go.mod", "db-service/go.sum", "./db-service/"]
COPY ["kafka-service/go.mod", "kafka-service/go.sum", "./kafka-service/"]

RUN go mod download