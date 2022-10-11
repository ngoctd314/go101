FROM golang:1.18-alpine

WORKDIR /app 

COPY main.go .

RUN go build -o ./run main.go

CMD ["./run"]