FROM golang:latest

WORKDIR WORKDIR /go/src/github.com/MicBun/go-job-list-api

COPY . .

RUN go get -d -v ./...

RUN go build -o go-job-list-api .

EXPOSE 8080

ENTRYPOINT ["./go-job-list-api"]