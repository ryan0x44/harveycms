FROM golang:1.6-alpine
COPY . /go/src/web
WORKDIR /go/src/web
RUN go build -o main .
CMD ["/go/src/web/main"]
EXPOSE 8080
