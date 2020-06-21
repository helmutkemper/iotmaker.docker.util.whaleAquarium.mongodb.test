FROM golang:alpine3.12 as builder

RUN mkdir /app
RUN chmod 700 /app

RUN apk add --no-cache git
RUN go get -u go.mongodb.org/mongo-driver; exit 0

COPY . /app

# import golang packages to be used inside image "scratch"
ARG CGO_ENABLED=0
RUN go build -o /app/main /app/main.go

FROM scratch

COPY --from=builder /app/ .

EXPOSE 8080
EXPOSE 27017

CMD ["/main"]
