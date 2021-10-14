FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build 
ENTRYPOINT [ "/app/hello" ]
