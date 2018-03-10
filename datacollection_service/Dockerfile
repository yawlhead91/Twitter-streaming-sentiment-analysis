FROM golang:latest
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /go/src/datacollection


CMD ["datacollection"]
