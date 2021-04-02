FROM golang:1.16

WORKDIR /go/src/app

ENV ACCOUNT_ID="nope"
ENV LICENSE_KEY="not bloody likely"

COPY . .

RUN go get -d -v ./...
RUN go generate
RUN go install -v ./...

CMD ["avoxi-challenge"]
