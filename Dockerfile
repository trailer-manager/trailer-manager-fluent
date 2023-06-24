FROM golang:1.20-alpine3.18 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY ./* .

RUN go mod download

RUN go build -o main

WORKDIR /app

RUN cp /app/main .

FROM scratch

COPY --from=builder /app/main .

ENTRYPOINT ["/app"]