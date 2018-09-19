FROM golang:1.11

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN ln -s /go/src/app/calc /go/src/app/vendor/calc
RUN ln -s /go/src/app/proto /go/src/app/vendor/proto
RUN ln -s /go/src/app/services /go/src/app/vendor/services

RUN go build -o /app/main .
CMD ["/app/main"]
