FROM golang:latest

RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go get github.com/robfig/cron

ENV TZ=Asia/Bangkok
WORKDIR /go/src/hello-monday
EXPOSE 6000

CMD ["go", "run", "src/main.go", "src/date.go"]