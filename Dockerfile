FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod init math

RUN go build -o math math.go

CMD ["./math"]