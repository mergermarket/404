FROM golang:1.12.4-alpine
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -v -d
RUN go install
CMD ["app"]
