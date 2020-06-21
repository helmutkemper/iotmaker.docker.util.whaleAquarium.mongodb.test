FROM golang:alpine3.12 as builder

RUN mkdir /app
RUN chmod 700 /app

COPY . /app

RUN apk add --no-cache git

RUN go get -u github.com/docker/docker
RUN go get -u github.com/helmutkemper/iotmaker.docker

RUN find . -name vendor -type d -print0|xargs -0 rm -rf --

# import golang packages to be used inside image "scratch"
ARG CGO_ENABLED=0
RUN go build -o /app/main /app/main.go

# VOLUME /app/static
EXPOSE 3000

CMD ["./main"]
