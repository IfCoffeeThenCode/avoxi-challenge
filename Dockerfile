FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get github.com/IfCoffeeThenCode/enumer@1.1.5
RUN go get -d -v ./...
RUN go generate
RUN go install -v ./...

CMD ["avoxi-challenge"]
