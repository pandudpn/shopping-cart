FROM golang:1.16-alpine3.13 AS builder

RUN apk update && apk add git

WORKDIR $GOPATH/src/shopping-cart

COPY . .

ENV GOSUMDB=off
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/shopping-cart cmd/main.go

FROM alpine:3.11

COPY --from=builder /go/bin/shopping-cart /go/bin/shopping-cart

RUN apk add --no-cache tzdata

ENTRYPOINT ["/go/bin/shopping-cart"]