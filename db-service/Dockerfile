FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o db-service

FROM alpine AS runner

COPY --from=builder /app/db-service /

CMD ["/db-service"]