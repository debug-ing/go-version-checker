FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build main.go

ENTRYPOINT ["/main"]