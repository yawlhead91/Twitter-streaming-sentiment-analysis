FROM golang:latest
RUN mkdir ./app
COPY  . /go/src/github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service
WORKDIR /go/src/github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service

RUN go get -d -v ./...
RUN go build

CMD app;