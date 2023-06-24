FROM golang:1.20-alpine3.18 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . "/usr/local/go/src/SiverPineValley/trailer-manager"
WORKDIR "/usr/local/go/src/SiverPineValley/trailer-manager"

RUN go mod download

RUN go build -o main

RUN mkdir /app
RUN mkdir /app/config

RUN cp ./main /app/main
RUN cp ./config/* /app/config

WORKDIR /app

FROM scratch

COPY --from=builder /app/main .

ENTRYPOINT ["./main --mode=dev"]