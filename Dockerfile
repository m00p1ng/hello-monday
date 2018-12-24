FROM golang:latest

WORKDIR /go/src/hello-monday
RUN go get github.com/line/line-bot-sdk-go/linebot
CMD ["go", "run", "main.go"]