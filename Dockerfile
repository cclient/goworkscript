FROM golang:1.8.0-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN go build src/main/apiserver.go
CMD ["/app/apiserver"]
