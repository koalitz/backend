FROM golang:alpine3.18 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /koalitz

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o /app/main cmd/main.go
ADD /configs /app/configs


FROM debian:buster-slim

ENV PROD 1

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/configs /app/configs

CMD ["./main"]