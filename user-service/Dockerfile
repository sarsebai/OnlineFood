FROM golang:1.24

WORKDIR /app

COPY . .

WORKDIR /app/user-service

RUN go mod download
RUN go build -o main .

CMD ["./main"]
