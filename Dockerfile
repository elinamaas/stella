FROM golang:1.9

RUN apt-get update && apt-get install -y cowsay

ADD . /go/src/github.com/elinamaas/stella

RUN go install github.com/elinamaas/stella/stella

CMD ["/go/bin/cowbot"]