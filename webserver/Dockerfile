# build docker binary
FROM golang:1.22.10-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /webserver cmd/webserver/main.go

EXPOSE 1234

# -D disables setting password
RUN adduser -D nonroot nonroot

RUN chown nonroot:nonroot /webserver

USER nonroot:nonroot

ENTRYPOINT ["/webserver"]