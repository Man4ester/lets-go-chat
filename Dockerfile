FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build cmd/lets-go-chat/main.go
CMD ["./main"]