# syntax=docker/dockerfile:1
##
## Build
##
FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod ./
# RUN go mod download

COPY *.go ./

RUN go build -o /go-headers

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-headers /go-headers

EXPOSE 8082

USER nonroot:nonroot

ENTRYPOINT ["/go-headers"]