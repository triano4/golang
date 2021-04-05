# FROM golang:1.15
# RUN  mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go build -o main .
# CMD ["/app/main"]

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]

# EXPOSE 8010