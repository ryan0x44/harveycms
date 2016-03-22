FROM golang:1.6-alpine
COPY . /go/src/api
WORKDIR /go/src/api
RUN go build -o main .
CMD ["/go/src/api/main"]
EXPOSE 8080
