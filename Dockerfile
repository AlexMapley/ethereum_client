FROM golang

ADD . /go/src/ethereum_client

RUN /bin/bash -c "export GOPATH=/go/src"
