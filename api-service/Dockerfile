FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o api-service ./cmd

FROM alpine AS runner

COPY --from=builder /app/api-service/ /api-service

CMD ["/api-service"]