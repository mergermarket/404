FROM golang:1.6-alpine as builder
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -v -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=builder /go/src/app/app /app
CMD ["/app"]
