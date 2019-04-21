FROM golang:1.11

RUN mkdir /go_test

WORKDIR /go_test

ADD . /go_test

#install chi
RUN go get github.com/go-chi/chi
RUN go get github.com/zephinzer/godev
RUN go install github.com/go-chi/chi

# RUN go build ./main.go

# CMD ["./main"]