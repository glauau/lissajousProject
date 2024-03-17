FROM golang:1.19

WORKDIR /go/src/app
COPY lissajous.go .

RUN go build -o lissajous lissajous.go

EXPOSE 8080

CMD ["./lissajous"]
