FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]