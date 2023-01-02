# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as build

RUN apk add --no-cache git
#ENV GOBIN /go/bin
#WORKDIR $GOPATH/app

WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# install all dependencies
#RUN go get ./...
WORKDIR /tmp/app/src
RUN go build -o ./docker-go-backend .

FROM alpine:3.9
COPY --from=build /tmp/app/src/docker-go-backend /app/docker-go-backend

EXPOSE 8080
CMD [ "/app/docker-go-backend" ]