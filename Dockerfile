FROM golang:latest

COPY . /go/src/github.com/flowup/docker-workshop

RUN go get github.com/flowup/docker-workshop/...
RUN go build -o /app/docker-service github.com/flowup/docker-workshop

RUN mkdir -p /app/data

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./docker-service"]