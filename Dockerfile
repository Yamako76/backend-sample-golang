FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

CMD ["./main"]
