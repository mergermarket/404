FROM golang:1.14.2-alpine
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -v -d
RUN go install
CMD ["app"]
