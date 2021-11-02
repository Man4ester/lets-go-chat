FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build cmd/rest/main.go
CMD ["./main"]