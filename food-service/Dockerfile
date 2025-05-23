FROM golang:1.24

WORKDIR /app

COPY . .

WORKDIR /app/food-service

RUN go mod download
RUN go build -o main .

CMD ["./main"]
