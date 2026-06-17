FROM golang:1.21-alpine AS builder

WORKDIR /app/cli

COPY cli/go.mod go.mod

RUN go mod download

COPY cli/main.go main.go

RUN go build -o monorepo main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cli/monorepo .

ENTRYPOINT ["./monorepo"]
